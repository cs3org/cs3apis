#!/bin/bash

set -ex  # Exit on error; debugging enabled.
set -o pipefail # Fail a pipe if any sub-command fails.

pushd /home/travis

PROTOBUF_VERSION=3.3.0
PROTOC_FILENAME=protoc-${PROTOBUF_VERSION}-linux-x86_64.zip
PROTOTOOL_VERSION=1.6.0

wget https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/${PROTOC_FILENAME}
unzip ${PROTOC_FILENAME}
sudo cp bin/protoc /usr/local/bin/protoc
protoc --version


# install prototool
curl -sSL \
  https://github.com/uber/prototool/releases/download/v${PROTOTOOL_VERSION}/prototool-Linux-x86_64.tar.gz | \
  sudo tar -C /usr/local --strip-components 1 -xz

which prototool

# install protolock
curl -sSL https://github.com/nilslice/protolock/releases/download/v0.12.0/protolock.20190327T205335Z.linux-amd64.tgz -o protolock.tgz && \ 
  tar -xzf protolock.tgz && cp protolock /usr/local/bin/

which protolock
popd
