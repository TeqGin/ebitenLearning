
all:genResource build

genResource:
	go-bindata -o src/resource/data.go resource/...

build:
	go build -o warGame ./src/listenKey/