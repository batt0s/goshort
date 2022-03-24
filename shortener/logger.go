package shortener

import (
	"log"
	"os"
)

// LOGGER Settings and Logger struct

type Logger struct {
	infoLogger   *log.Logger
	warnLogger   *log.Logger
	errorLogger  *log.Logger
	routerLogger *log.Logger
}

func (logger *Logger) InitLogger() {
	flags := log.LstdFlags | log.Lshortfile
	log.SetFlags(flags)
	infoLogger := log.New(os.Stdout, "INFO: ", flags)
	warnLogger := log.New(os.Stdout, "WARNING: ", flags)
	errorLogger := log.New(os.Stdout, "ERROR: ", flags)
	routerLogger := log.New(os.Stdout, "ROUTER: ", log.LstdFlags)
	logger.infoLogger = infoLogger
	logger.warnLogger = warnLogger
	logger.errorLogger = errorLogger
	logger.routerLogger = routerLogger
}

func (logger *Logger) Info(v ...interface{}) {
	logger.infoLogger.Println(v...)
}

func (logger *Logger) Warn(v ...interface{}) {
	logger.warnLogger.Println(v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.errorLogger.Println(v...)
}

func (logger *Logger) RouterLog(format string, v ...interface{}) {
	logger.routerLogger.Printf(format, v...)
}

var CustomLogger = Logger{}
