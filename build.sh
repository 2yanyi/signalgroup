#!/usr/bin/env bash

export GOPATH=$(pwd)

GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o bin/ddchat
