.PHONY: build
build:
	go build -o bin/clone ./cmd

binaries:
	GOOS=darwin GOARCH=amd64 go build -o bin/clone-amd64-darwin ./cmd
	GOOS=linux GOARCH=amd64 go build -o bin/clone-amd64-linux ./cmd
