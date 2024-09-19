package repeatrule

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
)

var log = slog.New(
	slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
)

func Verification(task models.Task) (string, error) {
	//проверяем поля запроса
	//проверка на наличие заголовка (обязательное поле)
	if task.Title == "" {
		log.Info("требуется заголовок задачи")
		return "", fmt.Errorf("требуется заголовок задачи")
	}
	//проверяем дату: если нет - то текущая дата, если есть то парсим дату
	var dateTask time.Time = time.Now()
	var err error
	if task.Date != "" {
		dateTask, err = time.Parse(models.DateFormat, task.Date) //парсим дату задачи в формате "20060102"
		if err != nil {
			log.Error("Ошибка при парсинге даты задачи", logger.Err(err))
			return "", fmt.Errorf("неверный формат даты задачи")
		}
	}

	//находим следующую дату задачи согласно правилу повторения
	now := time.Now()

	var nextDateTask string
	nextDateTask = now.Format(models.DateFormat) //дата будет текущей, если не будет вычислена новая

	switch {
	//если есть правило повторения и дата задачи в прошлом, то ищем следующую дату задачи
	case task.Repeat != "" && dateTask.Before(now.AddDate(0, 0, -1)):
		nextDateTask, err = NextDate(now, dateTask.Format(models.DateFormat), task.Repeat)
		if err != nil {
			log.Error("Ошибка при получении следующей даты задачи", logger.Err(err))
			return "", fmt.Errorf("ошибка при получении следующей даты задачи")
		}
	case dateTask.After(now.AddDate(0, 0, -1)): //Если дата задач в будущем
		nextDateTask = dateTask.Format(models.DateFormat)
	}

	return nextDateTask, nil
}
