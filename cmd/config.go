package main

import (
	"flag"
	"os"
)

//уровни логирования
const (
	envLocal = "local" //уровень по умолчанию
	envDev   = "dev"
	envProd  = "prod"
)

//флаги конфигурации
var FlagServerPort string //адрес порта
var FlagLevelLogger string //уровень логирования
// var FlagDatabaseDSN string      //наименование базы данных

//функция инициализации флагов
func ParseFlags() {

	defaultRunAddr := ":7540" //адрес порта по умолчанию
	defaultLevelLogger := "local" //уровень логирования по умолчанию

	flag.StringVar(&FlagServerPort, "a", defaultRunAddr, "port to run server") //инициализация флага адреса порта
	//flag.StringVar(&FlagDatabaseDSN, "d", defaultDatabaseDSN, "name database Postgres")
	flag.StringVar(&FlagLevelLogger, "l", defaultLevelLogger, "log level") //инициализация флага уровня логирования

	flag.Parse()

	if envRunAddr := os.Getenv("TODO_PORT"); envRunAddr != "" {
		FlagServerPort = envRunAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		FlagLevelLogger = envLogLevel
	}
	//envDatabaseDSN, ok := os.LookupEnv("DATABASE_URI")
	//if ok && envDatabaseDSN != "" {
	//	FlagDatabaseDSN = envDatabaseDSN
	//}
}
