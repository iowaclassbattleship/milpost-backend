.PHONY: build
build:
	@GOPATH=${PWD}/.gopath go build -o server main.go