package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct setup
type Config struct {
	WatchDir  string
	OutputDir string
	Encoder   string
	Preset    string
	Format    string
	Args      []string
}

// LoadConfig - used to import configuration
func LoadConfig() Config {
	var config Config
	cf, err := os.Open("/etc/merlin/config.json")
	defer cf.Close()
	if err != nil {
		log.Fatal("Cannot Open Config File!", err)
	}
	jp := json.NewDecoder(cf)
	jp.Decode(&config)
	switch config.Format {
	case "av_mp4":
		config.Format = ".m4v"
	case "av_mkv":
		config.Format = ".mkv"
	case "av_webm":
		config.Format = ".webm"
	}
	return config
}
