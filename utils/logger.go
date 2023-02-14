package utils

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func WriteLogging() *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.ErrorLevel)

	now := time.Now().UTC()
	fileName := "application-" + now.Format("YYYYMMDD") + ".log"
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	log.SetOutput(file)

	return log
}
