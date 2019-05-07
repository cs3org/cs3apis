# [CS3APIS](https://cernbox.github.io/cs3apis/) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Gitter chat](https://badges.gitter.im/cs3org/cs3apis.png)](https://gitter.im/cs3org/cs3apis) [![Build Status](https://travis-ci.org/cernbox/cs3apis.svg?branch=master)](https://travis-ci.org/cernbox/cs3apis)

**NOTE:** this repository in under heavy development
and not ready for public consumption.

CS3 stands for Cloud Storage Synchronisation and Sharing.

This repository contains the interface definitions of public
CS3 APIs that support both REST and gRPC protocols. You can also
use these definitions with open source tools to generate client
libraries, documentation, and other artifacts.

CS3 APIs follows Google API design guidelines, specially on error handling and naming convention.
You can read more about these guidelines at https://cloud.google.com/apis/design/.

This repository structure is very similar to https://github.com/googleapis/googleapis.

## Build it yourself
You need to have [Go](https://golang.org/doc/install), [git](https://git-scm.com/) and [make](https://en.wikipedia.org/wiki/Make_(software)) installed.

```
$ git clone https://github.com/cernbox/cs3apis
$ make deps # only supported on Darwin and Linux OS.
$ make
```

## Overview

CS3 APIs use [Protocol Buffers](https://github.com/google/protobuf)
version 3 (proto3) as their Interface Definition Language (IDL) to
define the API interface and the structure of the payload messages. The
same interface definition is used for both REST and RPC versions of the
API, which can be accessed over different wire protocols.

There are several ways of accessing CS3 APIs:

1.  Protocol Buffers over gRPC: You can access CS3 APIs published
in this repository through [GRPC](https://github.com/grpc), which is
a high-performance binary RPC protocol over HTTP/2. It offers many
useful features, including request/response multiplex and full-duplex
streaming.

2.  JSON over HTTP: You can access all CS3 APIs directly using JSON
over HTTP, using generated client libraries from the api definitions or 
or third-party API client libraries.

## Repository Structure

This repository uses a directory hierarchy that reflects the CS3
feature set. In general, every API has its own root
directory, and each major version of the API has its own subdirectory.
The proto package names exactly match the directory: this makes it
easy to locate the proto definitions and ensures that the generated
client libraries have idiomatic namespaces in most programming
languages. 

**NOTE:** The major version of an API is used to indicate breaking
change to the API.


## Official generated code

Go: [go-cs3apis](https://github.com/cs3org/go-cs3apis)

## Getting started

The following requirements are needed to compile the CS3 apis:

* Protobuf
* Prototool

### Installation

First we need to install the dependencies:

```
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
curl -sSL \
  https://github.com/nilslice/protolock/releases/download/v0.12.0/protolock.20190327T205335Z.linux-amd64.tgz | \
  sudo tar -C /usr/local --strip-components 1 -xz
```

### Compiling

From the root of the directory run:

```
make
```

## License

CS3 APIs are distributed under [Apache 2.0 license](https://github.com/cs3org/cs3apis/blob/master/LICENSE).
