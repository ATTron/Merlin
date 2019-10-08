# Merlin [![Go Report Card](https://goreportcard.com/badge/github.com/attron/merlin)](https://goreportcard.com/report/github.com/attron/merlin)
Merlin is a video encoder that automatically monitors for new video files and begins converting them via the HandBrakeCLI

## Setup
**Currently this project only supports Linux as it uses the inotify API for file system monitoring**

* Pull latest release
    * ``` git clone https://github.com/ATTron/merlin.git```

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

## Building for developers
**Developers wishing to contribute should ensure they are in a LINUX environment in order for proper functionality to be seen**
```
    make development
```  

## Contributing
[Contributions](https://github.com/attron/merlin/issues?q=is%3Aissue+is%3Aopen) are welcome, if you are interested please fork this repo and create a PR.

## License
MIT License, see [LICENSE](https://github.com/attron/merlin/blob/master/LICENSE)