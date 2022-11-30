BUILD_FOLDER=out
BINARY_NAME=asuswrt-api

all: clean build test

clean:
	go clean
	rm -rf ${BUILD_FOLDER}/*

build:
	go build -v -o ${BUILD_FOLDER}/ ./...

test:
	go test -v ./...
