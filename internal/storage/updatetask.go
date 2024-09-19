package storage

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
)

// UpdateTask обновляет задачу в базе данных
func (s *Storage) UpdateTask(ctx context.Context, task models.Task) error {
	op := "storage.UpdateTask"
	log := s.log.With(
		slog.String("op", op),
	)

	stmt, err := s.db.PrepareContext(ctx, `
		UPDATE scheduler
		SET date = ?, title = ?, comment = ?, repeat = ?
		WHERE id = ?
	`)
	if err != nil {
		log.Error("Ошибка при подготовке запроса на обновление задачи", logger.Err(err))
		return fmt.Errorf("ошибка при подготовке запроса на обновление задачи: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		log.Error("Ошибка при выполнении запроса на обновление задачи", logger.Err(err))
		return fmt.Errorf("ошибка при выполнении запроса на обновление задачи: %w", err)
	}

	//проверяем, что задача с таким ID существует и была обновлена	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error("Ошибка при получении количества измененных строк", logger.Err(err))
		return fmt.Errorf("ошибка при получении количества измененных строк: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("задача с ID %s не найдена", task.ID)
	}

	log.Info("Задача успешно обновлена")
	return nil
}