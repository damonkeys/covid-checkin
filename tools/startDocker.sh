#! /bin/bash

if [ -z "$1" ]
then
    echo -e "\nERROR: Missing command for starting... eg. './startDocker.sh dev build'"
    echo -e "  * First argument is the stage: dev or prod."
    echo -e "  * Second argument is optional."
    echo -e "         If you want to rebuild all code and images add the 'build' command."
    echo -e "         If you want to pull from our docker registry server add the 'pull' command.\n\n"
    exit
fi

cd ../docker

if [ ! -z "$2" ] && [ "$2" = 'build' ]
then
    ./buildAll.sh
fi

if [ ! -z "$2" ] && [ "$2" = 'pull' ]
then
    docker-compose -p chckr --env-file ./env/$1/.env pull
fi

docker-compose -p chckr --env-file ./env/$1/.env up -d --remove-orphans
