#!/bin/bash

export GOOS=windows
export GOARCH=amd64

echo $GOOS
echo $GOARCH


go build -o mycmd.exe
