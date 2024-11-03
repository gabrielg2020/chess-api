package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log Global logger
var Log *logrus.Logger

func init() {
	Log = logrus.New()

	Log.SetOutput(os.Stdout)

	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
		},
	})
	Log.SetLevel(logrus.DebugLevel)
}
