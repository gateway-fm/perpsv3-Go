package logger

import (
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
	"sync"
)

type logger struct {
	*logrus.Logger
}

var once sync.Once
var instance *logger

func Log() *logger {
	once.Do(func() {
		log := logrus.New()

		log.SetLevel(logrus.InfoLevel)

		log.SetOutput(os.Stdout)
		log.SetFormatter(&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %layer%: %msg%\n",
		})

		instance = &logger{log}
	})

	return instance
}
