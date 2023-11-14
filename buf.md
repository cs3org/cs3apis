# CS3APIs

This repository contains the interface definitions of public
CS3APIS that support the gRPC protocol.
You can also use these definitions with open source tools to generate client
libraries, documentation, and other artifacts.

CS3 APIs follows Google and Uber API design guidelines, specially on error handling and naming convention.
You can read more about these guidelines at https://cloud.google.com/apis/design/ and https://github.com/uber/prototool/blob/dev/style/README.md.

This repository structure is very similar to https://github.com/googleapis/googleapis.

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
