package logger

import (
	"log"
	"os"
)

type Loggers struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

var Log = &Loggers{
	InfoLog:  infoLog,
	ErrorLog: errorLog,
}
