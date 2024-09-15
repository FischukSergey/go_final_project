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

// SaverTask интерфейс для сохранения задачи
type SaveTasker interface {
	SaveTask(task models.SaveTask) (string, error)
}

// SaveTask api сохраняет задачу в базу данных
func SaveTask(log *slog.Logger, db SaveTasker) http.HandlerFunc {
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

		/*
		//проверяем поля запроса
		//проверка на наличие заголовка (обязательное поле)
		if task.Title == "" { 
			log.Info("требуется заголовок задачи")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.TaskResponse{Error: "Не указан заголовок задачи"})
			return
		}
		//проверяем дату: если нет - то текущая дата, если есть то парсим дату
		var dateTask time.Time = time.Now()
		var err error
		if task.Date != "" {
			dateTask, err = time.Parse("20060102", task.Date) //парсим дату задачи в формате "20060102"
			if err != nil {
				log.Error("Ошибка при парсинге даты задачи", logger.Err(err))
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, models.TaskResponse{Error: "Неверный формат даты задачи"})
				return
			}
		}

		//находим следующую дату задачи согласно правилу повторения
		now := time.Now()
		var nextDateTask string
		nextDateTask = now.Format("20060102") //дата будет текущей, если не будет вычислена новая
		//если есть правило повторения и дата задачи в прошлом, то ищем следующую дату задачи
		if task.Repeat != "" && dateTask.Before(now.AddDate(0, 0, -1)) {
			nextDateTask, err = repeatrule.NextDate(now, dateTask.Format("20060102"), task.Repeat)
			if err != nil {
				log.Error("Ошибка при получении следующей даты задачи", logger.Err(err))
				w.WriteHeader(http.StatusInternalServerError)
				render.JSON(w, r, models.TaskResponse{Error: "Ошибка при получении следующей даты задачи"})
				return
			}
		}
		*/

		nextDateTask, err := repeatrule.Verification(task)
		if err != nil {
			log.Error("Ошибка при верификации задачи", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.TaskResponse{Error: "Ошибка при верификации задачи"})
			return
		}

		//сохраняем задачу в базу данных
		id, err := db.SaveTask(models.SaveTask{
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
