package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLClient struct {
	DB *gorm.DB
}

func NewMysqlClient(host, username, password, database string, port int, config *gorm.Config) (*MySQLClient, error) {
	tz := "&loc=Asia%2FJakarta"
	if config == nil {
		config = &gorm.Config{}
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true%s", username, password, host, port, database, tz)
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		logrus.Error(fmt.Sprintf("Cannot connect to MySQL. %v", err))
		return nil, err
	}

	if db == nil {
		panic("missing database client")
	}

	return &MySQLClient{
		DB: db,
	}, nil
}
