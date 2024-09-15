package updatetask

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/FischukSergey/go_final_project/internal/storage"
	"github.com/go-chi/render"
)

type IUpdateTask interface {
	UpdateTask(ctx context.Context, task models.Task) error
}

func UpdateTask(log *slog.Logger, db *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Обновление задачи")

		var task models.SearchTask
		err := render.DecodeJSON(r.Body, &task)
		if err != nil {
			log.Error("Ошибка при декодировании обновляемой задачи", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}
		
		err = db.UpdateTask(r.Context(), task)
		if err != nil {
			log.Error("Ошибка при обновлении задачи", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, models.SearchTask{})
	}
}
