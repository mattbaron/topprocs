#!/bin/bash

OUTPUT_DIR=exe

mkdir -p $OUTPUT_DIR

#
# Build for current platform (i.e. macOS/Darwin)
#
go build -o $OUTPUT_DIR/topprocs.local cmd/main.go

#
# Build for linux/386
#
export GOOS=linux
export GOARCH=386

go build -o $OUTPUT_DIR/topprocs cmd/main.go

#
# Build for Windows/386
#
export GOOS=windows
export GOARACH=386

go build -o $OUTPUT_DIR/topprocs.exe cmd/main.go
