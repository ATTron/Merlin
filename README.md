# Merlin [![Go Report Card](https://goreportcard.com/badge/github.com/attron/merlin)](https://goreportcard.com/report/github.com/attron/merlin)
Merlin is a video encoder that automatically monitors for new video files and begins converting them via the HandBrakeCLI

## Setup
**Currently this project only supports Linux as it uses the inotify API for file system monitoring**

* Pull latest release or clone repo
    * ``` git clone https://github.com/attron/merlin.git```  
    * ``` wget https://github.com/ATTron/merlin/releases/download/1.x.x/merlin-linux.tar.gz```

* Install [HandBrakeCLI](https://handbrake.fr/downloads.php)
    * If you are on RHEL, CentOS, or Debian you can use the Makefile if you do not have HandBrakeCLI installed ```make install-handbrake```

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

* Start merlin  
**You may need to run merlin as sudo if you are getting a permission denied error**  
```
   make start 
```

## Building for development
**In order to accurately test the functionality of merlin you will need to run it on a LINUX machine with golang installed**
```
    make development
```  

## Contributing
[Contributions](https://github.com/attron/merlin/issues?q=is%3Aissue+is%3Aopen) welcome. If interested,fork this repo and submit a PR.

## License
MIT License, see [LICENSE](https://github.com/attron/merlin/blob/master/LICENSE)