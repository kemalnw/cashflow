package cmd

import (
	"github.com/kemalnw/cashflow/internal/base/app"
	"github.com/kemalnw/cashflow/internal/base/handler"
	"github.com/kemalnw/cashflow/pkg/db"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

var (
	appConfig   *app.Config
	baseHandler *handler.BaseHTTPHandler
	mysqlClient *db.MySQLClient
)

func initLog() {
	logrus.SetOutput(os.Stdout)
	if appConfig.IsProd() {
		// Only Warn and Error log for prod
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetLevel(logrus.WarnLevel)

		return
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	level := logrus.InfoLevel
	if appConfig.IsDebug() {
		level = logrus.DebugLevel
	}
	logrus.SetLevel(level)
}

func initDatabase() {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	gormConfig := &gorm.Config{}
	if appConfig.IsDebug() {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		)

		// Show sql executed in console to debug
		gormConfig = &gorm.Config{
			Logger: newLogger,
		}
	}

	mysqlClient, _ = db.NewMysqlClient(host, username, password, database, port, gormConfig)
}

func initInfrastructure() {
	appConfig = app.InitConfig()

	initLog()
	initDatabase()
}

func initHTTP() {
	initInfrastructure()

	baseHandler = handler.NewBaseHandler(appConfig, mysqlClient.DB)
}
