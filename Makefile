.PHONY: build all test clean

all: build

build: build-http-server

build-http-server:
	go build -o bin/envoy-builder src/cmd/envoy-builder/main.go

start:
	go run src/cmd/envoy-builder/main.go

swagger:
	swag init --pd -g ./src/cmd/envoy-builder/main.go
	
test:
	PROJECT_PWD=$(shell pwd) go test -race ./...

test-cover:
	go test -cover --count=1 ./...

cover-total:
	go test -cover --count=1 ./... -coverprofile cover.out
	go tool cover -func cover.out | grep total 

cov:
	PROJECT_PWD=$(shell pwd) go test -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

fmt:
	@gofmt -w $(shell find . -type f -name '*.go' -not -path './*_test.go')

fmt-check:
	@gofmt -l $(shell find . -type f -name '*.go' -not -path './*_test.go')

clean:
	@ro -fR bin
