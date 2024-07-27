package test

import (
	"compass_mini_api/internal/abstraction"
	"compass_mini_api/internal/config"
	db "compass_mini_api/pkg/database"
	"net/http"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func Init(t *testing.T) (*abstraction.Context, *gorm.DB) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
		}).Fatal("Load .env file error")
	}
	var PASSWORD string
	err = config.LoadForTest(PASSWORD)
	if err != nil {
		logrus.Fatal(err)
	}
	db.Init()
	conn, err := db.Connection("POSTGRES")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	require.NoError(t, err)
	e := echo.New()
	ctx := &abstraction.Context{
		Context: e.NewContext(&http.Request{}, nil),
		Auth: &abstraction.AuthContext{
			ID: 99,
		},
	}
	return ctx, conn
}
