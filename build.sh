#!/bin/bash
set -x o
set -e 
GOPATH=`go env GOPATH`
echo ${GOPATH}
sed "s|<GOPATH>|${GOPATH}|g" prototool-template.yaml > prototool.yaml
prototool compile
