all:build

genResource:
	go-bindata -pkg resource -o src/resource/data.go resource/...

build:genResource
ifeq ($(shell go env GOOS),windows)
	go build -o warGame.exe ./src/warGame/
else
	go build -o warGame_mac ./src/warGame/
endif

run:
ifeq ($(shell go env GOOS),windows)
	./warGame.exe
else
	./warGame_mac
endif

buildMac:
	SET CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	go build -o warGame_mac ./src/warGame/

buildWindows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o warGame.exe ./src/warGame/