pwd = $(shell pwd)
default: gen

gen:
	docker run -v ${pwd}:/root/cs3apis hugo cs3apis-build
	#docker run -v ${pwd}:/root/cs3apis cs3org/cs3apis cs3apis-build

clean:
	rm -rf build/
all: gen clean

