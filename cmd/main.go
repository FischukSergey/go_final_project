package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FischukSergey/go_final_project/internal/handlers/nextdate"
	"github.com/FischukSergey/go_final_project/internal/handlers/savetask"
	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/storage"

	"github.com/go-chi/chi"
)

func main() {
	ParseFlags()                                        //инициализируем флаги/переменные окружения конфигурации сервера
	log := setupLogger(FlagLevelLogger)                 //инициализируем логер с заданным уровнем
	db, err := storage.NewStorage(FlagDatabaseDSN, log) //инициализируем хранилище
	if err != nil {
		log.Error("Ошибка при подключении к базе данных", logger.Err(err))
		os.Exit(1)
	}
	defer db.Close()
	log.Info("База данных подключена", slog.String("database", FlagDatabaseDSN))

	//инициализируем роутер
	r := chi.NewRouter()
	root := "./web" //путь к статическим файлам
	//подключаем обработчик для статических файлов
	fileServer := http.FileServer(http.Dir(root))
	r.Handle("/*", fileServer)

	//подключаем обработчики для api
	r.Get("/api/nextdate", nextdate.NextDate(log))
	r.Post("/api/task", savetask.SaveTask(log, db))

	srv := &http.Server{ //инициализируем сервер
		Addr:         FlagServerPort,   //порт сервера
		Handler:      r,                //роутер
		ReadTimeout:  4 * time.Second,  //время на чтение запроса
		WriteTimeout: 4 * time.Second,  //время на запись ответа
		IdleTimeout:  30 * time.Second, //время на закрытие соединения
	}

	log.Info("Запуск сервера")
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
		log.Error("Ошибка при остановке сервера", logger.Err(err))
		return
	}
	log.Info("api server остановлен")
}

// функция инициализации логера
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
