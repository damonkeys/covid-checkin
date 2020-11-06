#! /bin/bash

echo -e "\nStarting goreleaser for building binary and docker image for $1"
echo -e "==========================================================================\n"
cd ../$1
goreleaser --snapshot --skip-validate --rm-dist
