#!/usr/bin/env bash
./buildEmbeddedFileServerVar.sh linux amd64

id=$(git rev-parse HEAD)
docker build -t $REGISTRY_SERVER/chckr/landing-$(basename $PWD):$id .
