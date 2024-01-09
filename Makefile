ifeq ($(shell go env GOOS),windows)
	warGameName = warGame.exe
	snakeGameName = snakeGame.exe
else ifeq ($(shell go env GOOS),linux)
	warGameName = warGame_linux
	snakeGameName = snakeGame_linux
else ifeq ($(shell go env GOOS),darwin)
	warGameName = warGame_mac
	snakeGameName = snakeGame_mac
endif


all:build

genSnakeGameResource:
	go-bindata -pkg resource -o src/resource/data.go resource/snake/...

genWarGameResource:
	go-bindata -pkg resource -o src/resource/data.go resource/war/...

build:
	@echo $(shell go env GOOS)
	make genWarGameResource && go build -o $(warGameName) ./src/warGame/
	make genSnakeGameResource && go build -o $(snakeGameName) ./src/snake

runWarGame:
	./$(warGameName)

runSnakeGame:
	./$(snakeGameName)




crossBuild:
ifeq ($(shell go env GOOS),windows)
	SET CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	go build -o warGame_mac ./src/warGame/
	go build -o snakeGame_mac ./src/snake
	SET GOOS=linux
	go build -o warGame_linux ./src/warGame/
	go build -o snakeGame_linux ./src/snake
else ifeq ($(shell go env GOOS),darwin)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o warGame.exe ./src/warGame/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o snakeGame.exe ./src/snake
else ifeq ($(shell go env GOOS),linux)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o warGame.exe ./src/warGame/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o snakeGame.exe ./src/snake
endif