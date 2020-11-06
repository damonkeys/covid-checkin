#! /bin/bash

if [ -z "$1" ]
then
    stage="dev"
else
    stage=$1
fi

curl --data-urlencode "config@../docker/kong/$stage/kong.yml" -X POST http://localhost:8001/config
