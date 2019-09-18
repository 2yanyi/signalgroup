#!/usr/bin/env bash
export GOPATH=$(pwd)

go build -ldflags '-w -s' -o bin/ddchat
