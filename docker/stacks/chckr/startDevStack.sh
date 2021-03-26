#! /bin/bash
cd $(dirname "$0")

source ./env/prod/urls.env
export DOMAIN_NAME=${DOMAIN_NAME}
export BASE_URL=${HTTP_PROTOCOL}://${DOMAIN_NAME}

docker stack deploy -c docker-compose.dev.yml --with-registry-auth chckr
