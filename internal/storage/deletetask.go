package storage

import (
	"context"
	"log/slog"

	"github.com/FischukSergey/go_final_project/internal/logger"
)

// DeleteTask удаляет задачу по id
func (s *Storage) DeleteTask(ctx context.Context, idTask int) error {
	op := "storage.DeleteTask"
	log := s.log.With(slog.String("op", op))

	stmt, err := s.db.PrepareContext(ctx, `DELETE FROM scheduler WHERE id = ?`)
	if err != nil {
		log.Error("Ошибка при подготовке запроса", logger.Err(err))
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, idTask)
	if err != nil {
		log.Error("Ошибка при выполнении запроса", logger.Err(err))
		return err
	}
	return nil
}