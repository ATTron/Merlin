# Merlin [![Go Report Card](https://goreportcard.com/badge/github.com/attron/merlin)](https://goreportcard.com/report/github.com/attron/merlin)
Merlin is a video encoder that automatically monitors for new video files and begins converting them via the HandBrakeCLI

## Setup
**Currently this project only supports Linux as it uses the inotify API for file system monitoring**

* Pull latest release or clone repo
    * ``` git clone https://github.com/attron/merlin.git```  
    * ``` wget https://github.com/ATTron/merlin/archive/1.x.x.zip```

* Install [HandBrakeCLI](https://handbrake.fr/downloads.php)

* Install and setup default configuration
```
    sudo make install
```

You can find a generated **config.json** in ``` /etc/merlin/ ```

The config.json file takes the following arguments:
```
   WatchDir  <- The directory for Merlin to watch
   OutputDir <- Directory that Merlin will store output files
   Encoder   <- The chosen encoder format (run HandBrakeCLI --help to see list of available encoder formats)
   Preset    <- The HandBrakeCLI presets (run HandBrakeCLI -z to see available video presets)
   Format    <- The output format for the encoded files (av_mp4, av_mkv, av_webm)
```

```
   make start 
```

## Building for development
**In order to accurately test the functionality of merlin you will need to run it on a LINUX machine**
```
    make development
```  

## Contributing
[Contributions](https://github.com/attron/merlin/issues?q=is%3Aissue+is%3Aopen) welcome. If interested,fork this repo and send submit a PR.

## License
MIT License, see [LICENSE](https://github.com/attron/merlin/blob/master/LICENSE)