package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/rjeczalik/notify"
)

var config Config

func main() {
	setupLogging()
	config, err := LoadConfig()
	if err != nil {
		return
	}

	c := make(chan notify.EventInfo, 1)
	done := make(chan bool)
	if err := notify.Watch(config.WatchDir, c, notify.InCloseWrite); err != nil {
		log.Fatal(err)
	}
	defer notify.Stop(c)

	go func() {
		for {
			select {
			case ei := <-c:
				go handleEvent(ei)
			}
		}
	}()

	<-done
}

func handleEvent(e notify.EventInfo) {
	// log.Printf("Event: %s", e.Event().String())
	ext := filepath.Ext(e.Path())
	done := make(chan bool, 1)
	if e.Event() == notify.InCloseWrite && filepath.Base(e.Path())[0:1] != "." {
		fn := filepath.Base(e.Path())
		log.Println("New File Detected:", e.Path())
		found := contains(ext)
		if found {
			go encode(e.Path(), done)
			d := <-done
			if d {
				log.Printf("Encoding Complete: %s", fn[:strings.LastIndex(fn, ".")]+config.Format)
			}
		} else {
			log.Printf("File type %s is not a supported video format\n", ext)
		}
	}
}
