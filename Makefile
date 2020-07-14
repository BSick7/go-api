SHELL := /bin/bash

.PHONY: test

test:
	go fmt ./...
	go test -v ./...
