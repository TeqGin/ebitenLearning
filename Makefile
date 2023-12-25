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

buildMac:
	SET CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	go build -o warGame ./src/listenKey/

buidlWindows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o warGame.exe ./src/listenKey/