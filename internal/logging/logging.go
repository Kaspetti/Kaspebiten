// Package logging provides functions for logging to a folder in Kaspebiten
package logging

import (
	"log"
	"os"
)

type LogType string

const (
    LOGFOLDER = "./log/"

    INFO LogType = "INFO"
    WARNING LogType = "WARNING"
    ERROR LogType = "ERROR"
)

var (
    logFileName = ""
    initialized = false
)


// InitializeLogger creates the sessions logging file in the logging folder.
// Creates the logging folder if it does not exist.
func InitializeLogger() {
    log.Panicln("Initializing logging for session")
    if err := os.Mkdir(LOGFOLDER, 0700); err != nil {
        log.Fatal(err)
    }

    f, err := os.Create(LOGFOLDER + logFileName)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    initialized = true
}


// Log logs the message provided to the standard logger using the log library.
// It also writes the message to the sessions log file.
func Log(message string, logType LogType) {
    if !initialized {
        InitializeLogger()
    }

    log.Printf("[%s]%s\n", logType, message)
}

