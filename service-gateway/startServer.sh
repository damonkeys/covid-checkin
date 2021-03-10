#!/bin/bash    

source .env
export SSL_ACTIVE=${SSL_ACTIVE}
export SERVER_PORT=${SERVER_PORT}
export ROUTES_CONFIG=${ROUTES_CONFIG}

./service-gateway
