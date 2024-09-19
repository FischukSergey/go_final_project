package main

import (
	"os"

	"github.com/FischukSergey/go_final_project/internal/models"
)

// уровни логирования
const (
	envLocal = "local" //уровень по умолчанию
	envDev   = "dev"
	envProd  = "prod"
)

// флаги конфигурации
var FlagServerPort string  //адрес порта
var FlagLevelLogger string //уровень логирования
var FlagDatabaseDSN string //наименование базы данных

// функция инициализации флагов
func ParseFlags() {
	//инициализация флагов по умолчанию
	defaultRunAddr := ":7540"                      //адрес порта по умолчанию
	defaultLevelLogger := "local"                  //уровень логирования по умолчанию
	defaultDatabaseDSN := "./storage/scheduler.db" //наименование базы данных по умолчанию

	//парсинг переменных окружения

	if envRunAddr := os.Getenv("TODO_PORT"); envRunAddr != "" {
		FlagServerPort = envRunAddr
	} else {
		FlagServerPort = defaultRunAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		FlagLevelLogger = envLogLevel
	} else {
		FlagLevelLogger = defaultLevelLogger
	}
	if envDatabaseDSN := os.Getenv("TODO_DBFILE"); envDatabaseDSN != "" {
		FlagDatabaseDSN = envDatabaseDSN
	} else {
		FlagDatabaseDSN = defaultDatabaseDSN
	}
	if envPassword := os.Getenv("TODO_PASSWORD"); envPassword != "" {
		models.Pass	 = envPassword
	} else {
		models.Pass = ""
	}
}
