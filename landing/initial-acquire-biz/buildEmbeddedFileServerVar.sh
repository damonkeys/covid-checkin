#!/usr/bin/env bash

echo "dont forget to make a call to:
'go get -u github.com/gobuffalo/packr/packr' and
'go get -u github.com/gobuffalo/packr' once before you run this tool."


env GOOS=$1 GOARCH=$2 packr
env GOOS=$1 GOARCH=$2 packr build

echo "server has been build - run it with ./${PWD##*/}"
