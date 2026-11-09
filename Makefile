pwd = $(shell pwd)
default: gen

gen:
	echo ${pwd}
	docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis-build:master cs3apis-build

clean:
	rm -rf build/
all: gen clean

docs:
	buf generate --template '{"version":"v1","plugins":[{"plugin":"doc","out":"docs","opt":"docs/html.tmpl,index.html","strategy":"all"}]}'
.PHONY: docs
