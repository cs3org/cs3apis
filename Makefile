.PHONY: build lint docs deps

pwd = $(shell pwd)
default: build

build:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis go run build.go -build-proto
python:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis go run build.go -build-python
go:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis go run build.go -build-go
js:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis go run build.go -build-js

all: build python go js
