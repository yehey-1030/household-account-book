package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}

func InitServerLogger() {
	logrus.SetOutput(os.Stderr)
	level := logrus.InfoLevel

	if v, ok := os.LookupEnv("LOG_LEVEL"); ok {
		level, _ = logrus.ParseLevel(v)
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(UTCFormatter{&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "req_time",
		},
		TimestampFormat: "2006-01-02T15:04:05.000Z",
	}})

}
