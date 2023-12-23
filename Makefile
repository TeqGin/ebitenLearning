all:build

genResource:
	go-bindata -pkg resource -o src/resource/data.go resource/...

build:genResource
ifeq ($(shell go env GOOS),windows)
	go build -o warGame.exe ./src/listenKey/
else
	go build -o warGame ./src/listenKey/
endif

run:
ifeq ($(shell go env GOOS),windows)
	./warGame.exe
else
	./warGame
endif