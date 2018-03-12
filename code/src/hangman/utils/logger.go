package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

const (
	SERVER_MODE = "SERVER_MODE"
	ENABLED     = "enabled"
)

func EnableServerMode() {
	os.Setenv(SERVER_MODE, ENABLED)
}

func Logger() *logrus.Logger {
	logger := logrus.New()
	if ENABLED == os.Getenv(SERVER_MODE) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	return logger
}
