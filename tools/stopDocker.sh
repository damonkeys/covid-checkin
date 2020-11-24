#! /bin/bash

if [ -z "$1" ]
then
    echo -e "\nERROR: Missing command for stopping... eg. './stopDocker.sh dev'"
    echo -e "  * First argument is the stage: dev or prod or other...\n\n"
    exit
fi

cd ../docker
docker-compose -p chckr --env-file ./env/$1/.env down
