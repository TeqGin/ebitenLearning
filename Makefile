all:build

genResource:
	go-bindata -pkg resource -o src/resource/data.go resource/...

build:genResource
ifeq ($(shell go env GOOS),windows)
	go build -o warGame.exe ./src/warGame/
	go build -o snakeGame.exe ./src/snake
else
	go build -o warGame_mac ./src/warGame/
	go build -o snakeGame_mac ./src/snake
endif

runWarGame:
ifeq ($(shell go env GOOS),windows)
	./warGame.exe
else
	./warGame_mac
endif

runSnakeGame:
ifeq ($(shell go env GOOS),windows)
	./snakeGame.exe
else
	./snakeGame_mac
endif

crossBuild:
ifeq ($(shell go env GOOS),windows)
	SET CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	go build -o warGame_mac ./src/warGame/
	go build -o snakeGame_mac ./src/snake
else
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o warGame.exe ./src/warGame/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o snakeGame.exe ./src/snake
endif