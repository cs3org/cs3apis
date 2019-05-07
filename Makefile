.PHONY: build lint docs
default: build lint

build:
	prototool compile
	protolock status

lint:
	prototool format -w
	prototool lint

docs:
	rm -rf docs && mkdir docs && protoc --doc_out=./docs --doc_opt=html,index.html cs3/*/*.proto cs3/*/*/*.proto 
