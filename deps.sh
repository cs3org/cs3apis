#!/bin/bash

set -xe  # Exit on error; debugging enabled.
set -o pipefail # Fail a pipe if any sub-command fails.

case "$OSTYPE" in
  darwin*)  OS="darwin" && PROTOOS="osx" ;; 
  linux*)   OS="linux" && PROTOOS="linux" ;;
  *)        echo "unknown os, build yourself: $OSTYPE" && exit 1 ;;
esac

echo "OS = ${OS}"
echo "PROTOOS= ${PROTOOS}"

PROTOBUF_VERSION=3.7.1
PROTOC_FILENAME=protoc-${PROTOBUF_VERSION}-${PROTOOS}-x86_64.zip
PROTOTOOL_VERSION=1.6.0

curl -sSL https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/${PROTOC_FILENAME} -o protoc.zip
unzip -o protoc.zip
sudo cp bin/protoc /usr/local/bin/protoc
protoc --version


# install prototool
curl -sSL \
  https://github.com/uber/prototool/releases/download/v${PROTOTOOL_VERSION}/prototool-${OS}-x86_64.tar.gz | \
  sudo tar -C /usr/local --strip-components 1 -xz

# install protolock
curl -sSL https://github.com/nilslice/protolock/releases/download/v0.12.0/protolock.20190327T205335Z.${OS}-amd64.tgz -o protolock.tgz
tar -xzf protolock.tgz
sudo cp protolock /usr/local/bin/

# install protoc-gen-doc to generate the documentation for the APIS.
curl -sSL https://github.com/pseudomuto/protoc-gen-doc/releases/download/v1.3.0/protoc-gen-doc-1.3.0.${OS}-amd64.go1.11.2.tar.gz -o protoc-gen-doc.tar.gz
tar xzfv protoc-gen-doc.tar.gz
sudo cp protoc-gen-doc-*/protoc-gen-doc /usr/local/bin/

which protoc
which prototool
which protolock
which protoc-gen-doc

protoc --version
prototool version
