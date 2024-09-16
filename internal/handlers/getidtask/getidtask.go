package getidtask

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
)

// IGetIDTask - интерфейс для получения задачи по id
type IGetIDTask interface {
	GetIDTask(ctx context.Context, id int) (models.SearchTask, error)
}

// GetIDTask - обработчик для получения задачи по id
func GetIDTask(log *slog.Logger, db IGetIDTask) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Получение задач по id")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		
		id := r.URL.Query().Get("id")
		if id == "" { //если id не передан, то возвращаем ошибку
			w.WriteHeader(http.StatusBadRequest)
			log.Error("Не указан id пользователя")
			render.JSON(w, r, models.ErrorResponse{Error: "Не указан id пользователя"})
			return
		}

		idTask, err := strconv.Atoi(id)
		if err != nil { //если парсинг не удался, то возвращаем ошибку
			w.WriteHeader(http.StatusBadRequest)
			log.Error("Неверный формат id пользователя", slog.String("id", id))
			render.JSON(w, r, models.ErrorResponse{Error: "Неверный формат id пользователя"})
			return
		}

		var task models.SearchTask

		task, err = db.GetIDTask(r.Context(), idTask) //получаем задачу по id
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
