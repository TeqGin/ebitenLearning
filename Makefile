ifeq ($(shell go env GOOS),windows)
	warGameName = warGame.exe
	snakeGameName = snakeGame.exe
	ooxxGameName = ooxx.exe
else ifeq ($(shell go env GOOS),linux)
	warGameName = warGame_linux
	snakeGameName = snakeGame_linux
	ooxxGameName = ooxx_linux
else ifeq ($(shell go env GOOS),darwin)
	warGameName = warGame_mac
	snakeGameName = snakeGame_mac
	ooxxGameName = ooxx_mac
endif


all:build

genSnakeGameResource:
	go-bindata -pkg resource -o src/resource/data.go resource/snake/...

genWarGameResource:
	go-bindata -pkg resource -o src/resource/data.go resource/war/...

genOOXXGameResource:
	go-bindata -pkg resource -o src/resource/data.go resource/ooxx/...

build:
	@echo $(shell go env GOOS)
	make genWarGameResource && go build -o $(warGameName) ./src/warGame/
	make genSnakeGameResource && go build -o $(snakeGameName) ./src/snake
	make genOOXXGameResource && go build -o $(ooxxGameName) ./src/ooxx

runWarGame:
	./$(warGameName)

runSnakeGame:
	./$(snakeGameName)

runOOXXGame:
	./$(ooxxGameName)


crossBuild:
ifeq ($(shell go env GOOS),windows)
	SET CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	go build -o warGame_mac ./src/warGame/
	go build -o snakeGame_mac ./src/snake
	go build -o ooxx_mac ./src/ooxx
	SET GOOS=linux
	go build -o warGame_linux ./src/warGame/
	go build -o snakeGame_linux ./src/snake
	go build -o ooxx_linux ./src/ooxx
else ifeq ($(shell go env GOOS),darwin)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o warGame.exe ./src/warGame/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o snakeGame.exe ./src/snake
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ooxx.exe ./src/ooxx
else ifeq ($(shell go env GOOS),linux)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o warGame.exe ./src/warGame/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o snakeGame.exe ./src/snake
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ooxx.exe ./src/ooxx
endif