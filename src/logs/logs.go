package logs

import (
	"log"
	"os"
)

var logFile *os.File

func InitializeLogs() {
	logFile, _ = os.Create("app.log")
	log.SetOutput(logFile)
}

func PostLog(mesType, message string) {
	log.Println(mesType + ": " + message)
}

func CloseLogs() {
	logFile.Close()
}
