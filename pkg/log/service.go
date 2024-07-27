package log

import (
	"compass_mini_api/internal/config"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nullseed/logruseq"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	INDEX_LOG_ERROR    = "log_error"
	INDEX_LOG_ACTIVITY = "log_activity"
	INDEX_LOG_LOGIN    = "log_login"
)

func Init() {
	logrus.AddHook(logruseq.NewSeqHook("http://localhost:5341"))
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var level logrus.Level = logrus.InfoLevel
	if config.Get().Logging.LogrusLevel != 0 {
		switch config.Get().Logging.LogrusLevel {
		case 1:
			level = 1
		case 2:
			level = 2
		case 3:
			level = 3
		case 4:
			level = 4
		case 5:
			level = 5
		case 6:
			level = 6
		}
	}
	logrus.SetLevel(level)

	timeStamp := time.Now().UTC().Add(time.Hour * 7).Format(time.DateTime)
	replacements := map[string]string{
		"-": "",
		" ": "_",
		":": "",
	}
	for oldChar, newChar := range replacements {
		timeStamp = strings.ReplaceAll(timeStamp, oldChar, newChar)
	}
	logFilePath := "../log/" + timeStamp + ".log"

	hook := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   true,
	}
	logrus.SetOutput(hook)
	fmt.Printf("write logs in files %s.log", timeStamp)
}

func InsertErrorLog(ctx context.Context, log *LogError) error {
	return nil
}

func InsertActivityLog(ctx context.Context, log *LogError) error {
	return nil
}

func InsertLoginLog(ctx context.Context, log *LogError) error {
	return nil
}

func LogruswriteError(id string, message string) error {
	logrus.WithFields(logrus.Fields{
		"id":    id,
		"error": message,
	}).Error("An error occured")
	return nil
}
