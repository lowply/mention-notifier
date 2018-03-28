default: run

SRC_FILES = $(wildcard src/*.go)

test:
	go test -v -parallel=4 ./...

run:
	go build -o main $(SRC_FILES)
	LOCAL=true ./main

build:
	GOOS=linux GOARCH=amd64 go build -o main $(SRC_FILES)
	zip main.zip main
	rm main
