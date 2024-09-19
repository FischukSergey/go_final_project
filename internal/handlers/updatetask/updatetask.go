package updatetask

import (
	"context"
	"log/slog"
	"net/http"

	repeatrule "github.com/FischukSergey/go_final_project/internal/lib"
	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
)

// IUpdateTask интерфейс для обновления задачи
type IUpdateTask interface {
	UpdateTask(ctx context.Context, task models.Task) error
}

func UpdateTask(log *slog.Logger, db IUpdateTask) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Обновление задачи")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		var task models.Task
		err := render.DecodeJSON(r.Body, &task)
		if err != nil {
			log.Error("Ошибка при декодировании обновляемой задачи", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}
		//проверяем задачу на корректность. 
		//передаем в функцию task, который содержит в себе поля: Title, Date, Repeat
		nextDateTask, err := repeatrule.Verification(models.Task{
			Title: task.Title,
			Date: task.Date,
			Repeat: task.Repeat,
		})
		if err != nil {
			log.Error("Ошибка при верификации задачи", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}
		//обновляем параметры задачи
		task.Date = nextDateTask //обновляем дату задачи
		err = db.UpdateTask(r.Context(), task)
		if err != nil {
			log.Error("Ошибка при обновлении задачи", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, models.ErrorResponse{})
	}
}
