# CS3 APIs

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

## Overview

CS3 APIs use [Protocol Buffers](https://github.com/google/protobuf)
version 3 (proto3) as their Interface Definition Language (IDL) to
define the API interface and the structure of the payload messages. The
same interface definition is used for both REST and RPC versions of the
API, which can be accessed over different wire protocols.

There are several ways of accessing CS3 APIs:

1.  JSON over HTTP: You can access all CS3 APIs directly using JSON
over HTTP, using generated client libraries from the api definitions or 
or third-party API client libraries.

2.  Protocol Buffers over gRPC: You can access CS3 APIs published
in this repository through [GRPC](https://github.com/grpc), which is
a high-performance binary RPC protocol over HTTP/2. It offers many
useful features, including request/response multiplex and full-duplex
streaming.


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


## Generate code
The code genereated in the `gen` folder is generated using `github.com/uber/prototool generate`.
