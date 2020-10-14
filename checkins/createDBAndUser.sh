#!/bin/bash

source .env
echo "this script creates the initial checkins database: ${DB_NAME} and user ${DB_USER}."
echo "currently this happens on the same mariadb server on which the rest of the system is running but could be done elsewhere"
echo "You should only run this once"

mysql -u root -p -h 127.0.0.1 -e "CREATE DATABASE ${DB_NAME}; \
                                 CREATE USER '${DB_USER}'@'%' IDENTIFIED BY '${DB_PASSWORD}'; \
                                 GRANT ALL PRIVILEGES ON ${DB_NAME}.* TO '${DB_USER}'@'%' IDENTIFIED BY '${DB_PASSWORD}'; \
                                 flush privileges;"
