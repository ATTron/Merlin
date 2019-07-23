# Merlin
Merlin is a video encoder that automatically monitors for new video files and begins converting them via the HandBrakeCLI

## SETUP
**Currently this project only supports Linux as it uses the inotify API for file system monitoring**

* Pull latest release
    * ``` wget https://github.com/ATTron/merlin/archive/1.0.0.zip```

* Install [HandBrakeCLI](https://handbrake.fr/downloads.php)

* Install and setup default configuration
```
    sudo make install
```

You can find a generated **config.json** in ``` > /etc/merlin/ ```

The config.json file takes the following arguments:
    * WatchDir  <- The directory for Merlin to watch
    * OutputDir <- Directory that Merlin will store output files
    * Encoder   <- The chosen encoder format (run HandBrakeCLI --help to see list of available encoder formats)
    * Preset    <- The HandBrakeCLI presets (run HandBrakeCLI -z to see available video presets)

* Run merlin with ``` make start ```

