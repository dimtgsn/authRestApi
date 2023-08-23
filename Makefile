.PHONY: build
build:
	go build -v ./cmd/restapi

.PHONY: start
start: build;./restapi



.DEFAULT_GOAL := build
