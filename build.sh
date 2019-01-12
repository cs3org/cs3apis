#!/bin/bash
set -x o
set -e 
GOPATH=`go env GOPATH`
echo ${GOPATH}
cp prototool-template.yaml prototool.yaml
sed -i "" "s|<GOPATH>|${GOPATH}|g" prototool.yaml
prototool compile
