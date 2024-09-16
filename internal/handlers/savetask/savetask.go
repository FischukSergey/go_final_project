package savetask

import (
	"log/slog"
	"net/http"

	//"time"

	repeatrule "github.com/FischukSergey/go_final_project/internal/lib"
	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"

	"github.com/go-chi/render"
)

// ISaveTasker интерфейс для сохранения задачи
type ISaveTasker interface {
	SaveTask(task models.SaveTask) (string, error)
}

// SaveTask api сохраняет задачу в базу данных
func SaveTask(log *slog.Logger, db ISaveTasker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("saving task started")
		defer log.Debug("saving task finished")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		task := models.Task{}

		if err := render.DecodeJSON(r.Body, &task); err != nil { //декодируем тело запроса в структуру Task
			log.Error("Ошибка при декодировании JSON", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.TaskResponse{Error: "Ошибка при декодировании JSON"})
			return
		}

		nextDateTask, err := repeatrule.Verification(task) //проверяем задачу
		if err != nil {
			log.Error("Ошибка при верификации задачи", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.TaskResponse{Error: "Ошибка при верификации задачи"})
			return
		}

		id, err := db.SaveTask(models.SaveTask{ //сохраняем задачу в базу данных
			Date:    nextDateTask,
			Title:   task.Title,
			Comment: task.Comment,
			Repeat:  task.Repeat,
		})
		if err != nil {
			log.Error("Ошибка при сохранении задачи", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, models.TaskResponse{Error: "Ошибка при сохранении задачи"})
			return
		}
		log.Info("Задача сохранена", slog.String("id", id))
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, models.TaskResponse{ID: id})
	}
}
