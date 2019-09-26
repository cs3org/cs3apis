.PHONY: build lint docs deps

default: build-proto

build-proto:
	go run build.go -build-proto
deps-proto:
	go run build.go -deps-proto

deps-go:
	go run build.go -deps-go
build-go:
	go run build.go -build-go
push-go:
	go run build.go -push-go

build-python:
	go run build.go -build-python
push-python:
	go run build.go -push-python
	
	
	
all-proto: deps-proto build-proto
all-go: deps-go build-go push-go
all-python: build-python push-python
all: all-proto all-go all-python
	
	
