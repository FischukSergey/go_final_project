package gettask

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
)

// IGetTasks интерфейс для получения задач
type IGetTasks interface {
	GetTasks(ctx context.Context, dateTask, search string) ([]models.SearchTask, error)
}

// GetTasks получает задачи из базы данных
func GetTasks(log *slog.Logger, db IGetTasks) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Получение задач")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		var search string   //строковая переменная для поиска
		var dateTask string //строковая переменная для даты

		search = r.URL.Query().Get("search")

		//пробуем парсить дату
		date, err := time.Parse("02.01.2006", search)
		if err == nil {
			dateTask = date.Format("20060102") //если получилось парсить, то присваиваем дату в нужном формате
			search = ""
		}

		var tasks []models.SearchTask
		log.Info("Получение задач", slog.String("dateTask", dateTask), slog.String("search", search))
		tasks, err = db.GetTasks(r.Context(), dateTask, search)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Ошибка при получении задач", logger.Err(err))
			render.JSON(w, r, models.SearchTasksResponse{Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, models.SearchTasksResponse{Tasks: tasks})
	}
}
