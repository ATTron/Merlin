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
}

// LoadConfig - used to import configuration
func LoadConfig() (Config, error) {
	var config Config
	cf, err := os.Open("/etc/merlin/config.json")
	if err != nil {
		log.Println("Cannot Open Config File!", err)

		// only create config if it does not already exist
		if os.IsNotExist(err) {
			createConfig()
		}

		return config, err
	}
	defer cf.Close()

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
	return config, nil
}

func createConfig() {
	// default configuration
	config := Config{
		WatchDir:  "/Users/encoder/Desktop",
		OutputDir: "/Users/enocder/",
		Encoder:   "x264",
		Preset:    "Fast 1080p30",
		Format:    "av_mkv",
	}

	// create /etc/merlin
	if _, err := os.Stat("/etc/merlin"); os.IsNotExist(err) {
		if err := os.Mkdir("/etc/merlin", os.ModePerm); err != nil {
			log.Println("Could not create \"/etc/merlin\"", err)
			return
		}
	}

	// create config file
	cf, err := os.Create("/etc/merlin/config.json")
	if err != nil {
		log.Println("Could not create config file!", err)
		return
	}

	// write config to file
	b, _ := json.MarshalIndent(config, "", "  ")
	if _, err := cf.Write(b); err != nil {
		log.Println("Could not write configuration to config file!", err)
		return
	}
	cf.Close()

	log.Println(
		"Created default configuration: \"/etc/merling/config.json\" - ",
		"Please modify it to fit your needs and relaunch Merlin",
	)
}
