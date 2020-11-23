#!/bin/bash         

if [ -z "$1" ]
then
    echo -e "\nERROR: Missing command for starting... eg. './startDocker.sh dev build'"
    echo -e "  * First argument is the stage: dev or prod."
    echo -e "  * Second argument is optional. If you want to rebuild all code and images add the 'build' command.\n\n"
    exit
fi

./stopDocker.sh $1
./startDocker.sh $1
