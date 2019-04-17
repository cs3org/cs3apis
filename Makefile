.PHONY: build
default: build lint

build:
	prototool compile
	protolock status

lint:
	prototool format -w
	prototool lint
