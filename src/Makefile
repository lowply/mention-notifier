default: run

test:
	go test -v -parallel=4 ./...

run:
	go build -o main
	LOCAL=true ./main

build:
	GOOS=linux GOARCH=amd64 go build -o main
	zip ../main.zip main
	rm main
