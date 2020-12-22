#! /bin/bash

if [ -z "$1" ]
then
    echo -e "\nERROR: Missing command for starting... eg. './startStack.sh dev build'"
    echo -e "  * First argument is the stage: dev or prod."
    echo -e "  * Second argument is optional."
    echo -e "         If you want to rebuild all code and images add the 'build' command."
    exit
fi

cd ../docker
source ./env/$1/.env
export STAGE=${STAGE}
export DOCKER_REGISTRY_SERVER=${DOCKER_REGISTRY_SERVER}
export DOMAIN_NAME=${DOMAIN_NAME}
export BASE_URL=${BASE_URL}

if [ ! -z "$2" ] && [ "$2" = 'build' ]
then
    ./buildAll.sh
fi

docker stack deploy -c docker-compose.yml chckr
