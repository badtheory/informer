package informer

import (
	"errors"
	"github.com/creasty/defaults"
	)

var informer Logger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default informer level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	//InstanceZapLogger will be used to create Zap instance for the logger
	InstanceZapLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the Logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	EnableConsole     bool `default:"true"`
	ConsoleJSONFormat bool `default:"false"`
	ConsoleLevel      string `default:"informer.Debug"`
	EnableFile        bool `default:"true"`
	FileJSONFormat    bool `default:"true"`
	FileLevel         string `default:"informer.Debug"`
	FileLocation      string `default:"log.log"`
}

//NewLogger returns an instance of Logger
func NewLogger(config *Configuration, loggerInstance int) error {

	if err := defaults.Set(config); err != nil {
		panic(err)
	}

	if loggerInstance == InstanceZapLogger {
		logger, err := newZapLogger(config)
		if err != nil {
			return err
		}
		informer = logger
		return nil
	}
	return errInvalidLoggerInstance
}

func Debugf(format string, args ...interface{}) {
	informer.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	informer.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	informer.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	informer.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	informer.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	informer.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return informer.WithFields(keyValues)
}
