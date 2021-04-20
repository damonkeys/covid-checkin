#! /bin/bash
cd $(dirname "$0")

source ./env/staging/urls.env
export DOMAIN_NAME=${DOMAIN_NAME}
export BASE_URL=${HTTP_PROTOCOL}://${DOMAIN_NAME}

source ./env/staging/dbchckr.env
export DB_CHCKR_HOST=${DB_CHCKR_HOST}
export DB_CHCKR_NAME=${DB_CHCKR_NAME}
export DB_CHCKR_USER=${DB_CHCKR_USER}
export DB_CHCKR_PASSWORD=${DB_CHCKR_PASSWORD}
export DB_CHCKR_ROOT_PASSWORD=${DB_CHCKR_ROOT_PASSWORD}

source ./env/staging/dbcheckins.env
export DB_CHECKINS_HOST=${DB_CHECKINS_HOST}
export DB_CHECKINS_NAME=${DB_CHECKINS_NAME}
export DB_CHECKINS_USER=${DB_CHECKINS_USER}
export DB_CHECKINS_PASSWORD=${DB_CHECKINS_PASSWORD}
export DB_CHECKINS_ROOT_PASSWORD=${DB_CHECKINS_ROOT_PASSWORD}

docker stack deploy -c docker-compose.staging.yml --with-registry-auth chckr
