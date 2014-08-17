package main

import (
	log "github.com/bambocher/go-logging"
)

func main() {
	log.SetLevel(log.TRACE)
	log.SetFormat("[{asctime}][{levelname}][{filename}:{lineno}] {message}\n")
	log.SetDateFormat("2006-01-02 15:04:05")

	log.Trace("Trace message.")
	log.Debug("Debug message.")
	log.Info("Informational message.")
	log.Notice("Notice message.")
	log.Warning("Warning message.")
	log.Error("Error message.")
	//	log.Critical("Critical message.")
	//	log.Alert("Alert message.")
	//  log.Panic("Panic message.")
}
