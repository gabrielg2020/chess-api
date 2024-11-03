package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

// Log Global logger
var Log *logrus.Logger

func init() {
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
	Log.SetLevel(logrus.DebugLevel)
}
