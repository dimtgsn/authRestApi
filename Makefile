.PHONY: build
build:
	go build -v ./cmd/restapi

.PHONY: build
start: build;./restapi



.DEFAULT_GOAL := build
