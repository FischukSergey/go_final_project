package storage

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/FischukSergey/go_final_project/internal/logger"
	_ "github.com/mattn/go-sqlite3"
)

// структура для хранения данных
type Storage struct {
	db *sql.DB
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

	return &Storage{db: db}, nil
}

// функция закрытия базы данных
func (s *Storage) Close() error {
	return s.db.Close()
}
