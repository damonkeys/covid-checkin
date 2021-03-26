#! /bin/bash
cd $(dirname "$0")

docker stack deploy -c docker-compose.prod.yml --with-registry-auth www
