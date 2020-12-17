#!/bin/bash    

source .env
export SERVER_PORT_SSL=${SERVER_PORT_SSL}
export ROUTES_CONFIG=${ROUTES_CONFIG}

./service-gateway
