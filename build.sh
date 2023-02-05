#!/bin/bash

#
# Build for current platform (i.e. macOS/Darwin)
#
go build -o topprocs.local cmd/topprocs/main.go

#
# Build for linux/386
#
export GOOS=linux
export GOARCH=386

go build -o topprocs cmd/topprocs/main.go

#
# Build for Windows/386
#
export GOOS=windows
export GOARACH=386

go build -o topprocs.exe cmd/topprocs/main.go
