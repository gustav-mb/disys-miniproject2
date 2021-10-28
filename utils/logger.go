package utils

import (
	fmt "fmt"
	"log"
	"os"
)

type Logger struct {
	Log *log.Logger
}

func NewLogger(filename string) *Logger {
	f, err := os.OpenFile("..\\logs\\"+filename+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	//defer(f.Close())

	return &Logger{
		Log: log.New(f, filename+" ", log.LstdFlags),
	}
}

func (l *Logger) PrintlnBoth(text string) {
	l.Logln(text)
	log.Println(text)
}

func (l *Logger) Println(text string) {
	l.Logln(text)
	fmt.Println(text)
}

func (l *Logger) Logln(text string) {
	l.Log.Println(text)
}

func (l *Logger) logf(text string, v ...interface{}) {
	l.Log.Printf(text, v...)
}

func (l *Logger) logFatalf(text string, v ...interface{}) {
	l.Log.Fatalf(text, v...)
}
