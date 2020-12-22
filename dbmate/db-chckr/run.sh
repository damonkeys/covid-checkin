#!/bin/bash
source ../.env

#needed: ./run.sh up or ./run.sh migrate
# generally speaking: ./run.sh <databasename> <dbmate_command>
dburl=mysql://$DB_CHCKR_USER:$DB_CHCKR_PASSWORD@$DB_CHCKR_HOST:3306/$DB_CHCKR_NAME dbmate -e dburl $1
