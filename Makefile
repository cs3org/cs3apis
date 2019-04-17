.PHONY: build
default: build lint

build:
	prototool compile

lint:
	prototool format -w
	prototool lint
