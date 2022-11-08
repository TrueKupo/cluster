package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger = nil

	global_logger *logrus.Entry = nil
)

func LogModule(module string) *logrus.Entry {
	return global_logger.WithFields(logrus.Fields{
		"module": module,
	})
}

func InitLogger(log_level string, log_format string, log_file string, stderr bool, args logrus.Fields) {
	log = logrus.New()

	logLevel, err := logrus.ParseLevel(log_level)
	if err != nil {
		logLevel = logrus.ErrorLevel
	}

	log.SetLevel(logLevel)

	if log_format == "json" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		})
	}

	if log_file == "stderr" || stderr == true {
		log.SetOutput(os.Stderr)
	} else {
		file, err := os.OpenFile(log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Warn("Failed to log to file, using default stderr")
		}
	}

	global_logger = log.WithFields(
		logrus.Fields(args),
	)

	return
}
