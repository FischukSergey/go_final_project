package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// Storage структура для хранения данных
type Storage struct {
	db  *sql.DB
	log *slog.Logger
}

// NewStorage функция инициализации хранилища
func NewStorage(storagePath string, log *slog.Logger) (*Storage, error) {

	db, err := sql.Open("sqlite3", storagePath) //инициализируем базу данных
	if err != nil {
		log.Error("Ошибка при создании таблицы пользователей", logger.Err(err))
		return nil, err
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		"date" CHAR(8) NOT NULL,
		title TEXT,
		comment TEXT,
		repeat CHAR(128)
	    );`,
	)
	if err != nil {
		log.Error("Ошибка при создании шаблона SQL запроса", logger.Err(err))
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error("Ошибка закрытия шаблона SQL запроса", logger.Err(err))
		}
	}(stmt)

	_, err = stmt.Exec()
	if err != nil {
		log.Error("Ошибка при создании таблицы", logger.Err(err))
		return nil, err
	}

	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);")
	if err != nil {
		log.Error("Ошибка создания индекса", logger.Err(err))
		return nil, err
	}

	return &Storage{db: db, log: log}, nil
}

// SaveTask функция сохранения задачи
func (s *Storage) SaveTask(task models.Task) (string, error) {
	ctx := context.Background()
	const op = "storage.SaveTask"
	log := s.log.With(
		slog.String("op", op),
	)

	//подготавливаем запрос на сохранение задачи
	stmt, err := s.db.PrepareContext(ctx, `
	INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?);
	`)
	if err != nil {
		log.Error("Ошибка при сохранении задачи", logger.Err(err))
		return "", err
	}
	defer stmt.Close()

	//выполняем запрос на сохранение задачи
	result, err := stmt.ExecContext(ctx, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		log.Error("Ошибка при сохранении задачи", logger.Err(err))
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error("Ошибка при получении id сохраненной задачи", logger.Err(err))
		return "", err
	}
	log.Info("Задача сохранена", slog.String("id", fmt.Sprintf("%d", id)))
	//возвращаем id сохраненной задачи
	return fmt.Sprintf("%d", id), nil
}

// Close функция закрытия базы данных
func (s *Storage) Close() error {
	return s.db.Close()
}
