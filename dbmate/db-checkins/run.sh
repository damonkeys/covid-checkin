#!/bin/bash
source ../.env

#needed: ./run.sh up or ./run.sh migrate
# generally speaking: ./run.sh <databasename> <dbmate_command>
dburl=mysql://$DB_CHECKINS_USER:$DB_CHECKINS_PASSWORD@$DB_CHECKINS_HOST:3306/$DB_CHECKINS_NAME dbmate -e dburl $1
