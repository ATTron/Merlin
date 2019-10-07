package main

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

var allowedTypes = []string{"mp4", "mkv", "mp4", "m4a", "m4v", "mov", "webm", "flv", "f4v"}

func contains(s string) bool {
	for _, v := range allowedTypes {
		if "."+v == s {
			return true
		}
	}
	return false
}

// setup logging for both syslog and stdout
func setupLogging() {
	f, err := os.Create("Merlin-Encoder.log")
	if err != nil {
		log.Fatal("Cannot create log file!")
	}
	lw, err := syslog.New(syslog.LOG_NOTICE, "merlin")
	mw := io.MultiWriter(os.Stdout, lw, f)
	if err == nil {
		log.SetOutput(mw)
	}

	log.Println("Setup of logging complete")
}
