package logger

import (
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const filename = "/var/log/transact.log"

//GetLogger logger factory
func GetLogger(file string) *log.Logger {
	e, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("error opening file", err)
	}

	logger := log.New(e, file+" ", log.Ldate|log.Ltime)
	logger.SetOutput(&lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})

	return logger
}
