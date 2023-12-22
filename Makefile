
all:genResource build

genResource:
	go-bindata -pkg resource -o src/resource/data.go resource/...

build:
	go build -o warGame ./src/listenKey/

run:
	./warGame