#!/usr/bin/env bash
./buildEmbeddedFileServerVar.sh linux amd64
docker build -t chckr/landing-$(basename $PWD) -t ${{ secrets.REGISTRY_SERVER }}/chckr/landing-$(basename $PWD) .
