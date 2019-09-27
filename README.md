# [CS3APIS](https://cernbox.github.io/cs3apis/) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Gitter chat](https://badges.gitter.im/cs3org/cs3apis.svg)](https://gitter.im/cs3org/cs3apis) [![Build Status](https://cloud.drone.io/api/badges/cs3org/cs3apis/status.svg)](https://cloud.drone.io/cs3org/cs3apis)


The CS3APIS connect Cloud Storage and Applications Providers.

## API Documentation
https://cs3org.github.io/cs3apis/

## Build it yourself
You need to have Docker installed.

```
$ git clone https://github.com/cernbox/cs3apis
$ make build 
$ make go # generate go code
```

## Overview

This repository contains the interface definitions of public
CS3 APIs that support the gRPC protocol.
You can also use these definitions with open source tools to generate client
libraries, documentation, and other artifacts.

CS3 APIs follows Google API design guidelines, specially on error handling and naming convention.
You can read more about these guidelines at https://cloud.google.com/apis/design/.

This repository structure is very similar to https://github.com/googleapis/googleapis.

CS3 APIs use [Protocol Buffers](https://github.com/google/protobuf)
version 3 (proto3) as their Interface Definition Language (IDL) to
define the API interface and the structure of the payload messages. The
same interface definition is used for both REST and RPC versions of the
API, which can be accessed over different wire protocols.

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


## License

CS3 APIs are distributed under [Apache 2.0 license](https://github.com/cs3org/cs3apis/blob/master/LICENSE).
