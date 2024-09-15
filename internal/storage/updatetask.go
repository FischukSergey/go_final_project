package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
)

func (s *Storage) UpdateTask(ctx context.Context, task models.SearchTask) error {
	op := "storage.UpdateTask"
	log := s.log.With(
		slog.String("op", op),
	)
	
	stmt, err := s.db.PrepareContext(ctx, `
		UPDATE scheduler
		SET date = $2, title = $3, comment = $4, repeat = $5
		WHERE id = $1
	`)
	if err != nil {
		log.Error("Ошибка при подготовке запроса на обновление задачи", logger.Err(err))
		return fmt.Errorf("ошибка при подготовке запроса на обновление задачи: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, task.ID, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		log.Error("Ошибка при выполнении запроса на обновление задачи", logger.Err(err))
		return fmt.Errorf("ошибка при выполнении запроса на обновление задачи: %w", err)
	}
	log.Info("Задача успешно обновлена")
	return nil
}