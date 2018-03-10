package utils

import (
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	logger := logrus.New()
	return logger
}
