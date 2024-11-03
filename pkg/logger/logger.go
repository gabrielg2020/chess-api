package logger

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Log Global logger
var Log *logrus.Logger

func init() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Error loading .env file")
	}

	Log = logrus.New()

	Log.SetOutput(os.Stdout)

	Log.SetReportCaller(true)

	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			//funcName := path.Base(f.File)
			fileName := fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
			return fileName, ""
		},
		FullTimestamp: true,
	})

	// Set log level
	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		levelStr = "info" // default log level
	}

	level, err := logrus.ParseLevel(strings.ToLower(levelStr))
	if err != nil {
		Log.Warn("Error setting log level")
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)
}
