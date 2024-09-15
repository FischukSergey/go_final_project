package getidtask

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
)

type IGetIDTask interface {
	GetIDTask(ctx context.Context, id int) (models.SearchTask, error)
}

func GetIDTask(log *slog.Logger, db IGetIDTask) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Получение задач по id")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Error("Не указан id пользователя")
			render.JSON(w, r, models.ErrorResponse{Error: "Не указан id пользователя"})
			return
		}

		idTask, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error("Неверный формат id пользователя", slog.String("id", id))
			render.JSON(w, r, models.ErrorResponse{Error: "Неверный формат id пользователя"})
			return
		}

		var task models.SearchTask

		task, err = db.GetIDTask(r.Context(), idTask)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Ошибка при получении задачи по id", slog.String("id", id))
			render.JSON(w, r, models.ErrorResponse{Error: "Ошибка при получении задачи по id"})
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, task)
	}
}
