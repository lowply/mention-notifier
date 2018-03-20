default: run

SRC_FILES = $(wildcard *.go)

test:
	go test -v -parallel=4 .

run:
	LOCAL=true go run $(SRC_FILES)

build:
	GOOS=linux GOARCH=amd64 go build -o main $(SRC_FILES)
	zip main.zip main
	rm main
