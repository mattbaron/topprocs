#!/bin/bash

go build -o topprocs cmd/topprocs/main.go

export GOOS=linux
export GOARCH=386

go build -o topprocs.linux cmd/topprocs/main.go

export GOOS=windows
export GOARACH=386

go build -o topprocs.exe cmd/topprocs/main.go
