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
var singleInstance *counter

type counter struct {
	val int
}

func getCounter() *counter {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &counter{}
		}
	}
	return singleInstance
}

// logger example
// this example uses sync.Once which is a synchronization mechanism
// it guarantees that the code will be executed only once during the execution of the program
var once sync.Once
var logger *logrus.Logger

func getLogger() *logrus.Logger {
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

func RunSingleton() {
	for i := 0; i < 30; i++ {
		go getCounter()
		go getLogger()
	}
}
