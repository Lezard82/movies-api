package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrus.InfoLevel)
	})

	env := os.Getenv("APP_ENV")
	if env == "" {
		panic("APP_ENV is not defined in environment")
	}

	if env == "production" {
		logger.SetLevel(logrus.WarnLevel)
	} else {
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger
}
