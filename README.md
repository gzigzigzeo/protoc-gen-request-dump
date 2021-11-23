# protoc-request-dump

`protoc` plugin which dumps a request to the binary file. Used to test `protoc` generators.

## Installation

```
go install github.com/gzigzigzeo/protoc-gen-request-dump
```

## Usage

Will create `request.bin` in the current directory:

```
protoc --plugin=protoc-gen-request-dump --request-dump_out=request.bin:. example.proto
```