package logger

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type logger struct {
	*logrus.Logger
}

var once sync.Once
var instance *logger

func Log() *logger {
	once.Do(func() {
		log := logrus.New()

		levelEnv := os.Getenv("LOG_LEVEL")

		var logLevel logrus.Level

		switch levelEnv {
		case "trace":
			logLevel = logrus.TraceLevel
		case "debug":
			logLevel = logrus.DebugLevel
		case "info":
			logLevel = logrus.InfoLevel
		case "warn":
			logLevel = logrus.WarnLevel
		case "error":
			logLevel = logrus.ErrorLevel
		case "fatal":
			logLevel = logrus.FatalLevel
		case "panic":
			logLevel = logrus.PanicLevel
		default:
			logLevel = logrus.InfoLevel
		}

		log.SetLevel(logLevel)

		log.SetOutput(os.Stdout)
		log.SetFormatter(&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - Perpsv3-%layer%: %msg%\n",
		})

		instance = &logger{log}
	})

	return instance
}
