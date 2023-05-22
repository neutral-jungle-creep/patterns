package creational

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"sync"
)

// counter example
// this example uses mutex and checks singleInstance for nil
// second check needed if more than one goroutine passed the first
var lock = &sync.Mutex{}
var singleInstance *сounter

type сounter struct {
	val int
}

func GetCounter() *сounter {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &сounter{}
		}
	}

	return singleInstance
}

// logger example
// this example uses sync.Once which is a synchronization mechanism
// it guarantees that the code will be executed only once during the execution of the program
var once sync.Once
var logger *logrus.Logger

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger = logrus.New()
		logger.SetReportCaller(true)

		logger.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				fileName := path.Base(frame.File)
				return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", fileName, frame.Line)
			},
		}
		logger.SetLevel(logrus.InfoLevel)
	})
	return logger
}
