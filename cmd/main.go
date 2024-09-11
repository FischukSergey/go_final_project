package main

import (
	"context"
	"github.com/FischukSergey/go_final_project/internal/logger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
)

func main() {
	ParseFlags()                        //инициализируем флаги/переменные окружения конфигурации сервера
	log := setupLogger(FlagLevelLogger) //инициализируем логер с заданным уровнем

	r := chi.NewRouter()
	root := "./web"

	fileServer := http.FileServer(http.Dir(root))
	r.Handle("/*", fileServer) //обработчик для статических файлов

	srv := &http.Server{ //запускаем сервер
		Addr:         FlagServerPort,
		Handler:      r,
		ReadTimeout:  4 * time.Second,
		WriteTimeout: 4 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Info("Initializing server")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("Ошибка при запуске сервера", logger.Err(err))
		}
	}()
	log.Info("Server started", slog.String("address", srv.Addr))

	//Остановка процессов
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done //ждем сигнал прерывания
	//останавливаем сервер на прием новых запросов и дорабатываем принятые
	log.Info("Server stopping", slog.String("address", srv.Addr))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", logger.Err(err))
		return
	}
	log.Info("api server stopped")
}

//функция инициализации логера
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
