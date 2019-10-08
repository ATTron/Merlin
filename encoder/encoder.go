package main

import (
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func encode(s string, d chan bool) {
	fn := filepath.Base(s)
	oname := fn[:strings.LastIndex(fn, ".")] + config.Format
	log.Printf("Encoding Start: %s", s)
	cmd := exec.Command("HandBrakeCLI", "-i", s, "-o", config.OutputDir+oname, "-f", config.Format, "-Z", config.Preset, "--decomb=bob")
	if err := cmd.Start(); err != nil {
		log.Println("Error converting video")
	}
	err := cmd.Wait()
	if err != nil {
		log.Println("Conversion failed with code:", err)
	}
	d <- true
}
