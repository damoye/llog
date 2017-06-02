package llog

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stderr, "INFO: ", log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERRO: ", log.LstdFlags|log.Lshortfile)
)

// Info ...
func Info(v ...interface{}) {
	infoLogger.Output(2, fmt.Sprint(v...))
}

// Error ...
func Error(v ...interface{}) {
	errorLogger.Output(2, fmt.Sprint(v...))
}
