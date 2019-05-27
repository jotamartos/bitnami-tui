#!/bin/bash

mkdir /tmp/btui
GOOS=windows GOARCH=amd64 go build -o /tmp/btui/btui-windows-x64.exe
GOOS=darwin GOARCH=amd64 go build -o /tmp/btui/btui-osx-x64
GOOS=linux GOARCH=amd64 go build -o /tmp/btui/btui-linux-x64.run
