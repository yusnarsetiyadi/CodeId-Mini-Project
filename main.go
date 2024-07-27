package main

import (
	"compass_mini_api/internal/config"
	"compass_mini_api/internal/factory"
	"compass_mini_api/internal/http"
	"compass_mini_api/internal/middleware"
	db "compass_mini_api/pkg/database"
	"compass_mini_api/pkg/log"
	"compass_mini_api/pkg/ngrok"
	"context"
	httpNet "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var PASSWORD string

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
		}).Fatal("Load .env file error")
	}

	err = config.Load(PASSWORD)
	if err != nil {
		logrus.Fatal(err)
	}

	os.MkdirAll("../employeephoto", os.ModePerm)
	os.MkdirAll("../log", os.ModePerm)
}

// @title compass_mini_api
// @version 1.0.0
// @description This is a doc for compass_mini_api.

func main() {

	PORT := config.Get().Server.Port
	if PORT == "0" {
		PORT = "8080"
	}

	log.Init()
	db.Init()
	e := echo.New()
	f := factory.NewFactory()
	middleware.Init(e, f)
	http.Init(e, f)

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		runNgrok := true
		addr := ""
		if runNgrok {
			listener := ngrok.Run()
			e.Listener = listener
			addr = "/"
		} else {
			addr = ":" + PORT
		}
		err := e.Start(addr)
		if err != nil {
			if err != httpNet.ErrServerClosed {
				logrus.Fatal(err)
			}
		}
	}()

	<-ch

	logrus.Println("Shutting down server...")
	cancel()

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	e.Shutdown(ctx2)
	logrus.Println("Server gracefully stopped")
}
