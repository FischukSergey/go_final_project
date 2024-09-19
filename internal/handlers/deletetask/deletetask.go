// Package deletetask содержит обработчик для удаления задачи
package deletetask

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
)

// IDeleteTask интерфейс для удаления задачи
type IDeleteTask interface {
	DeleteTask(ctx context.Context, idTask int) error
}

// DeleteTask обработчик для удаления задачи
func DeleteTask(log *slog.Logger, db IDeleteTask) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		log.Info("API: Удаление задачи")

		id := r.URL.Query().Get("id")
		if id == "" { //если id не передан, то возвращаем ошибку
			log.Error("API: Не передан id задачи")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.ErrorResponse{Error: "Не передан id задачи"})
			return
		}

		idTask, err := strconv.Atoi(id)
		if err != nil { //если id не число, то возвращаем ошибку
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.ErrorResponse{Error: "Некорректный id задачи"})
			return
		}

		err = db.DeleteTask(r.Context(), idTask) //удаляем задачу
		if err != nil {
			log.Error("Ошибка при удалении задачи", slog.String("id", id), logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)	
		render.JSON(w, r, models.ErrorResponse{})
	}
}
