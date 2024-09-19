package storage

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
)

// GetIDTask получает задачу по id	
func (s *Storage) GetIDTask(ctx context.Context, idTask int) (models.Task, error) {
	op := "storage.GetIDTask"
	log := s.log.With(
		slog.String("op", op),
	)
	log.Info("Получение задачи по id", slog.Int("id task", idTask))

	stmt, err := s.db.PrepareContext(ctx, `
	SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?;
	`)
	if err != nil {
		log.Error("Ошибка при подготовке запроса", logger.Err(err))
		return models.Task{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, idTask)

	var task models.Task
	err = row.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("Задача не найдена", slog.Int("id task", idTask))
			return models.Task{}, err
		}
		log.Error("Ошибка при сканировании строки", logger.Err(err))
		return models.Task{}, err
	}

	return task, nil
}
