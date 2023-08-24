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

		log.SetLevel(logrus.InfoLevel)

		log.SetOutput(os.Stdout)
		log.SetFormatter(&easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - Perpsv3-%layer%: %msg%\n",
		})

		instance = &logger{log}
	})

	return instance
}
