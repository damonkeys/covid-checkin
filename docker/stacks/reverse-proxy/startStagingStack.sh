#! /bin/bash
cd $(dirname "$0")

docker stack deploy -c docker-compose.staging.yml --with-registry-auth proxy
