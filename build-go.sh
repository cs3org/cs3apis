#!/bin/bash

BRANCH=`git branch | grep \* | cut -d ' ' -f2`
GOPATH=`go env GOPATH`
PATH=$PATH:$GOPATH/bin
rm -rf build/go-cs3apis &&  mkdir build
cd build && git clone https://github.com/cs3org/go-cs3apis && cd go-cs3apis && git checkout -b ${BRANCH} && cd ../..
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
prototool generate
# change import paths
egrep -lR 'github.com/cs3org/go-cs3apis/' ./build/go-cs3apis  | xargs sed -i 's|github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/|github.com/cs3org/go-cs3apis/cs3/|g'
