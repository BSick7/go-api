SHELL := /bin/bash

.PHONY: all
all: tools dep test

tools:
	go get -u github.com/golang/dep/cmd/dep

dep:
	dep ensure

test:
	go test -v ./...
