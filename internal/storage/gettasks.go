package storage

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
)

func (s *Storage) GetTasks(ctx context.Context, dateTask, search string) ([]models.Task, error) {
	op := "storage.GetTasks"
	log := s.log.With(
		slog.String("op", op),
	)
	log.Info("Получение задач", slog.String("dateTask", dateTask), slog.String("search", search))
	var err error
	var rows *sql.Rows

	//выбираем запрос в зависимости от того, что задано
	switch {
	case dateTask != "": //если дата задана
		log.Info("Получение задач по дате", slog.String("dateTask", dateTask))
		stmt, err := s.db.PrepareContext(ctx, `
	SELECT * FROM scheduler WHERE date = ? LIMIT ?;
	`)
		if err != nil {
			log.Error("Ошибка при подготовке запроса на получение задач по дате", logger.Err(err))
			return nil, err
		}

		rows, err = stmt.QueryContext(ctx, dateTask, models.LimitTasks)
		if err != nil {
			log.Error("Ошибка при выполнении запроса на получение задач", logger.Err(err))
			return nil, err
		}
		defer rows.Close()

	case search != "": //если поиск по названию
		log.Info("Получение задач по названию", slog.String("search", search))
		stmt, err := s.db.PrepareContext(ctx, `
	SELECT * FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?;
	`)
		if err != nil {
			log.Error("Ошибка при подготовке запроса на получение задач по названию", logger.Err(err))
			return nil, err
		}

		searchsql := "%" + search + "%"
		rows, err = stmt.QueryContext(ctx, searchsql, searchsql, models.LimitTasks)
		if err != nil {
			log.Error("Ошибка при выполнении запроса на получение задач", logger.Err(err))
			return nil, err
		}
		defer rows.Close()

	default: //если не задано ничего, то выбираем все задачи
		log.Info("Получение всех задач")
		stmt, err := s.db.PrepareContext(ctx, `
	SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?;
	`)
		if err != nil {
			log.Error("Ошибка при подготовке запроса на получение всех задач", logger.Err(err))
			return nil, err
		}

		rows, err = stmt.QueryContext(ctx, models.LimitTasks)
		if err != nil {
			log.Error("Ошибка при выполнении запроса на получение задач", logger.Err(err))
			return nil, err
		}
		defer rows.Close()
	}
	//создаем массив для хранения задач
	var tasks = make([]models.Task, 0, models.LimitTasks)

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			log.Error("Ошибка при сканировании задачи", logger.Err(err))
			return nil, err
		}
		tasks = append(tasks, task)
	}
	err = rows.Err()
	if err != nil {
		log.Error("Ошибка завершения сканирования задач", logger.Err(err))
		return nil, err
	}

	return tasks, nil
}
