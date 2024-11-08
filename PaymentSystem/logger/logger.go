package logger

import (
	"fmt"
	"log"
	"os"
)

var logFile *os.File

// init initializes the logger file
func init() {
	var err error
	logFile, err = os.OpenFile("transactions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
}

// Log records a message to the log
func Log(message string) {
	defer logFile.Sync() // Ensure data is written to file
	log.Println(message)
	fmt.Println(message)
}
