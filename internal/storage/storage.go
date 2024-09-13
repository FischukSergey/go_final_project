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

// структура для хранения данных
type Storage struct {
	db *sql.DB
	log *slog.Logger
}

// функция инициализации хранилища
func NewStorage(storagePath string, log *slog.Logger) (*Storage, error) {

	db, err := sql.Open("sqlite3", storagePath) //инициализируем базу данных
	if err != nil {
		log.Error("Ошибка при создании таблицы пользователей", logger.Err(err))
		return nil, err
	}
	ctx := context.Background()
	stmt, err := db.PrepareContext(ctx, `
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		date CHAR(8),
		title TEXT,
		comment TEXT,
		repeat CHAR(128));
	CREATE INDEX IF NOT EXISTS idx_date ON scheduler (date);`)
	if err != nil {
		log.Error("Ошибка при создании таблицы пользователей", logger.Err(err))
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx)
	if err != nil {
		log.Error("Ошибка при создании таблицы пользователей", logger.Err(err))
		return nil, err
	}	

	return &Storage{db: db, log: log}, nil
}

//SaveTask функция сохранения задачи
func (s *Storage) SaveTask(task models.SaveTask) (string, error) {
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
