.PHONY: build lint docs deps

pwd = $(shell pwd)
default: build

pull:
	docker pull cs3org/cs3apis

build: pull
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-proto
	# restore ownership of the `docs` folder as docker runs as root
	chown -R `ls -ld . | awk '{print $$3 ":" $$4}'` docs
python: pull
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-python
go: pull
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-go
js: pull
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-js
node: pull
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build -build-node
clean:
	rm -rf build/

all: build python go js node
