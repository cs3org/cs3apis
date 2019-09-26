.PHONY: build lint docs deps

default: build

build:
	docker run cs3org/cs3apis go run build.go -build-proto
