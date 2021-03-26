#! /bin/bash
cd $(dirname "$0")

docker stack deploy -c docker-compose.dev.yml --with-registry-auth proxy
