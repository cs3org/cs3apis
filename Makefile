.PHONY: build lint docs deps
default: build lint docs
build-all: deps build lint docs

build:
	prototool compile
	protolock status

lint:
	prototool format -w
	prototool lint
	go run tools/check-license/check-license.go

docs:
	rm -rf docs && mkdir docs && protoc --doc_out=./docs --doc_opt=html,index.html cs3/*/*.proto cs3/*/*/*.proto 

deps:
	cp deps.sh /tmp && cd /tmp && ./deps.sh

build-go: deps build
	./build-go.sh

clean:
	rm -rf build
