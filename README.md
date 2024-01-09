# War Game
  - golang 1.20.12
  - engin: ebiten

## How to install go-bindata
  - go install github.com/go-bindata/go-bindata/...@latest


## some errors in linux
  1. Package 'alsa', required by 'virtual:world', not found
    `apt-get install libasound2-dev`
  2. /usr/bin/ld: cannot find -lXxf86vm: No such file or directory
    `apt-get install libxxf86vm-dev`
