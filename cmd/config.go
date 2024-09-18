package main

import (
	//"flag"
	"os"
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

	// flag.StringVar(&FlagServerPort, "a", defaultRunAddr, "port to run server")        //инициализация флага адреса порта
	// flag.StringVar(&FlagDatabaseDSN, "d", defaultDatabaseDSN, "name database SQLite") //инициализация флага наименования базы данных
	// flag.StringVar(&FlagLevelLogger, "l", defaultLevelLogger, "log level")            //инициализация флага уровня логирования

	//парсинг флагов	
	//flag.Parse()

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
}
