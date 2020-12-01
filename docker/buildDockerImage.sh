#! /bin/bash

echo -e "\nStarting building binary and docker image for $1"
echo -e "======================================================================\n"
cd ../$1
env GOOS=linux GOARCH=amd64 go build -o $1
docker build -t chckr/$1 -t ${{ secrets.REGISTRY_SERVER }}/chckr/$1 .
