package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	//ファイルの読み書き、追記、作成を指定
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
