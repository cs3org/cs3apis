[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Gitter chat](https://badges.gitter.im/cs3org/cs3apis.svg)](https://gitter.im/cs3org/cs3apis) [![Build Status](https://github.com/cs3org/cs3apis/actions/workflows/build-and-publish.yml/badge.svg)](https://github.com/cs3org/cs3apis/actions/workflows/build-and-publish.yml)
================

<img width="100px" src="https://raw.githubusercontent.com/cs3org/logos/master/cs3org/cs3org.png"/>

The CS3APIS connect Storage and Applications Providers.

## API Documentation
https://buf.build/cs3org-buf/cs3apis

## Officialy compiled libraries
The libraries for different languages are compiled from the protobuf definitions in this repo.
When a commit to master is made the CI takes care to create a new version of the library in the following languages.
Please note that the versioning used in the libraries below differs from language to language, however they point to the 
same source commit. This is due to the way the different package managers handle package versions.

* Go: https://github.com/cs3org/go-cs3apis
* Python: https://github.com/cs3org/python-cs3apis
* Javascript: https://github.com/cs3org/js-cs3apis (to be used from Web applications - frontend usage only)
* NodeJS: https://github.com/cs3org/node-cs3apis (to be used from NodeJS applications - backend)

## Repository packages
* Go: https://pkg.go.dev/github.com/cs3org/go-cs3apis
* Python: https://pypi.org/project/cs3apis/
* Javascript: https://www.npmjs.com/package/@cs3org/cs3apis
* NodeJS: https://www.npmjs.com/package/@cs3org/node-cs3apis


## Local compilation

You need to have Docker installed. The artifacts will be available under the build directory.

```
$ git clone https://github.com/cs3org/cs3apis
$ cd cs3apis
$ make
```
The `build` folder will be generated. It will be owned by `root`.

See the Makefile to find all the available build options.

## Overview

This repository contains the interface definitions of public
CS3APIS that support the gRPC protocol.
You can also use these definitions with open source tools to generate client
libraries, documentation, and other artifacts.

CS3 APIs use [Protocol Buffers](https://github.com/google/protobuf)
version 3 (proto3) as their Interface Definition Language (IDL) to
define the API interface and the structure of the payload messages.

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
