#!/bin/bash

source checkin/.env
export SERVER_PORT=${SERVER_PORT}
export DB_HOST=${DB_HOST}
export DB_NAME=${DB_NAME}
export DB_USER=${DB_USER}
export DB_PASSWORD=${DB_PASSWORD}
export SESSION_SECRET=${SESSION_SECRET}
echo $DB_HOST
./checkins
