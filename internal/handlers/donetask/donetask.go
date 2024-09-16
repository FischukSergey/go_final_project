package donetask

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	repeatrule "github.com/FischukSergey/go_final_project/internal/lib"
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
)

type IDoneTask interface {
	DeleteTask(ctx context.Context, idTask int) error
	GetIDTask(ctx context.Context, idTask int) (models.SearchTask, error)
	UpdateTask(ctx context.Context, task models.SearchTask) error
}

func DoneTask(log *slog.Logger, db IDoneTask) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("API: Завершение задачи")
		log.With(slog.String("method", r.Method), slog.String("path", r.URL.Path))

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//получаем id задачи
		id := r.URL.Query().Get("id")
		if id == "" { //если id не передан, то возвращаем ошибку
			log.Error("API: Не передан id задачи")
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.ErrorResponse{Error: "Не передан id задачи"})
			return
		}

		idTask, err := strconv.Atoi(id) //парсим id в int
		if err != nil {                 //если парсинг не удался, то возвращаем ошибку
			w.WriteHeader(http.StatusBadRequest)
			log.Error("Неверный формат id пользователя", slog.String("id", id))
			render.JSON(w, r, models.ErrorResponse{Error: "Неверный формат id пользователя"})
			return
		}

		var task models.SearchTask
		//получаем задачу по id
		task, err = db.GetIDTask(r.Context(), idTask)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Ошибка при получении задачи по id", slog.String("id", id))
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}

		// проверяем, что repeat не пустой
		switch {

		case task.Repeat != "": //если задача повторяется, то обновляем дату задачи
			log.Info("Задача повторяется", slog.String("repeat", task.Repeat))
			// nextDateTask, err := repeatrule.Verification(models.Task{ //получаем следующую дату задачи
			// 	Title:  task.Title,
			// 	Date:   task.Date,
			// 	Repeat: task.Repeat,
			// })
			nextDateTask, err := repeatrule.NextDate(time.Now(), task.Date, task.Repeat)
			fmt.Println(task,nextDateTask)
			if err != nil {
				log.Error("Ошибка при получении следующей даты задачи", slog.String("repeat", task.Repeat))
				w.WriteHeader(http.StatusInternalServerError)
				render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
				return
			}
			log.Info("Следующая дата задачи", slog.String("nextDateTask", nextDateTask))
			task.Date = nextDateTask
			err = db.UpdateTask(r.Context(), task) //обновляем задачу
			if err != nil {
				log.Error("Ошибка при обновлении задачи", slog.String("repeat", task.Repeat))
				w.WriteHeader(http.StatusInternalServerError)
				render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			render.JSON(w, r, models.ErrorResponse{})
			return

		default: //если задача не повторяется, то удаляем задачу
			log.Info("Задача не повторяется, удаляем задачу")
			err = db.DeleteTask(r.Context(), idTask)
			if err != nil {
				log.Error("Ошибка при удалении задачи", slog.String("repeat", task.Repeat))
				w.WriteHeader(http.StatusInternalServerError)
				render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
				return
			}
			w.WriteHeader(http.StatusOK)
			render.JSON(w, r, models.ErrorResponse{})
			return
		}
	}
}
