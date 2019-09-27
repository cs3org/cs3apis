.PHONY: build lint docs deps

pwd = $(shell pwd)
default: build

build:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-proto
python:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-python
go:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-go
js:
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-js
clean:
	rm -rf build/

all: build python go js
