#! /bin/bash

echo -e "\nStart building binary and docker image for $1"
echo -e "======================================================================\n"
cd ../$1

id=$(git rev-parse HEAD)
env GOOS=linux GOARCH=amd64 go build -o $1
docker build --tag ${{ secrets.REGISTRY_SERVER }}/chckr/$1:$id .
