package logger

import "log"

func Info(v ...any) {
	log.SetPrefix("INFO ")
	log.Println(v...)
}

func Error(v ...any) {
	log.SetPrefix("ERROR ")
	log.Println(v...)
}
