#!/bin/bash         

if [ -z "$1" ]
then
    echo -e "\nERROR: Missing command for starting... eg. './startDocker.sh dev'"
    echo -e "  * First argument is the stage: dev or prod or other."
    exit
fi

./stopStack.sh
./startStack.sh $1
