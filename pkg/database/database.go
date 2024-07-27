package database

import (
	"compass_mini_api/internal/config"
	"errors"
	"fmt"
	"strings"

	// "github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var dbConnections map[string]*gorm.DB

func Init() {

	dbConfigurations := map[string]Db{
		"POSTGRES": &dbPostgreSQL{
			db: db{
				Host: config.Get().DB.DbHost,
				User: config.Get().DB.DbUser,
				Pass: config.Get().DB.DbPass,
				Port: config.Get().DB.DbPort,
				Name: config.Get().DB.DbName,
			},
			SslMode: config.Get().DB.DbSslmode,
			Tz:      config.Get().DB.DbTz,
		},
	}

	dbConnections = make(map[string]*gorm.DB)
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (*gorm.DB, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}
