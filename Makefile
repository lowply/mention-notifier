default: run

SRC_FILES = $(wildcard *.go)

test:
	go test -v -parallel=4 .

run:
	go run $(SRC_FILES)

build:
	GOOS=linux GOARCH=amd64 go build -o main $(SRC_FILES)
	zip main.zip main
