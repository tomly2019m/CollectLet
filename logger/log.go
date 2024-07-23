package logger

import (
	"CollectLet/util"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Logger interface {
	Init()
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

type GlobalLogger struct {
	Logger
	Path string `yaml:"path"`
	wg   sync.WaitGroup
}

var logger *GlobalLogger

func GetLogger() *GlobalLogger {
	if logger == nil {
		logger = &GlobalLogger{}
		logger.Init()
	}
	return logger
}

func (l *GlobalLogger) Init() {
	byteValue, err := os.ReadFile("./config/logger.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(byteValue, l)
	if _, err := os.Stat(l.Path); err != nil {
		file, err := os.Create(l.Path)
		if err != nil {
			fmt.Println(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(file)
	}
}

func (l *GlobalLogger) logWith(logLevel string, format string, args ...interface{}) {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		file, err := os.OpenFile(l.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
		}
		_, err = fmt.Fprintf(file, util.GetCurrentTime()+" "+logLevel+": "+format+"\n", args...)
		if err != nil {
			fmt.Println(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(file)
	}()
}

func (l *GlobalLogger) WaitForDone() {
	l.wg.Wait()
}

func (l *GlobalLogger) Debug(format string, args ...interface{}) {
	logger.logWith("DEBUG", format, args...)
}

func (l *GlobalLogger) Info(format string, args ...interface{}) {
	logger.logWith("INFO", format, args...)
}

func (l *GlobalLogger) Warn(format string, args ...interface{}) {
	logger.logWith("WARN", format, args...)
}

func (l *GlobalLogger) Error(format string, args ...interface{}) {
	logger.logWith("ERROR", format, args...)
}

func (l *GlobalLogger) Fatal(format string, args ...interface{}) {
	logger.logWith("FATAL", format, args...)
}
