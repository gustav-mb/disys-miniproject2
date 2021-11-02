package utils

import (
	"log"
	"os"
)

const logPath = "..\\logs\\"

type Logger struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
}

// InfoPrintln Prints both to the log file and to the console.
// Equals to calling Logger.InfoLogger.Println() and a normal log.println()
func (l *Logger) InfoPrintln(v ...interface{}) {
	l.InfoLogger.Println(v...)
	log.Println(v...)
}

// InfoPrintf Prints a formated string to both the log file and to the console.
// Equals to calling Logger.InfoLogger.Printf() and a normal log.printf()
func (l *Logger) InfoPrintf(text string, v ...interface{}) {
	l.InfoLogger.Printf(text, v...)
	log.Printf(text, v...)
}

func (l *Logger) WarningPrintln(v ...interface{}) {
	l.WarningLogger.Println(v...)
	log.Println(v...)
}

func (l *Logger) ErrorFatalf(text string, v ...interface{}) {
	l.ErrorLogger.Printf(text, v...)
	log.Fatalf(text, v...)
}

func (l *Logger) ErrorPrintf(text string, v ...interface{}) {
	l.ErrorLogger.Printf(text, v...)
	log.Printf(text, v...)
}

func NewLogger(filename string) *Logger {
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		_ = os.Mkdir(logPath, os.ModeDir)
	}

	f, err := os.OpenFile(logPath+filename+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	return &Logger{
		InfoLogger:    log.New(f, "INFO: ", log.LstdFlags),
		WarningLogger: log.New(f, "WARNING: ", log.LstdFlags),
		ErrorLogger:   log.New(f, "ERROR: ", log.LstdFlags|log.Lshortfile),
	}
}
