default: run

SRC_FILES = $(wildcard *.go)

test:
	go test -v -parallel=4 .

run:
	go run $(SRC_FILES)

clean:
	rm -rf bin dist

lambda:
	GOOS=linux GOARCH=amd64 go build -o main $(SRC_FILES)
	zip main.zip main
	rm main

build: clean
	mkdir bin dist
	gox -osarch="darwin/amd64" \
		-osarch="linux/amd64" \
		-output="bin/{{.OS}}_{{.Arch}}/{{.Dir}}"
	zip -j dist/mention-notifier_darwin_amd64.zip bin/darwin_amd64/mention-notifier
	zip -j dist/mention-notifier_linux_amd64.zip bin/linux_amd64/mention-notifier
