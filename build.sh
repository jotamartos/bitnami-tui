#!/bin/bash

mkdir /tmp/bnhelper
GOOS=windows GOARCH=amd64 go build -o /tmp/bnhelper/bnhelper-windows-x64.exe
GOOS=darwin GOARCH=amd64 go build -o /tmp/bnhelper/bnhelper-osx-x64
GOOS=linux GOARCH=amd64 go build -o /tmp/bnhelper/bnhelper-linux-x64.run
