#!/bin/bash
source ../.env
#needed: ./run.sh up or ./run.sh migrate
# generally speaking: ./run.sh <databasename> <dbmate_command>
dburl=mysql://$DB_USER:$DB_PASSWORD@$DB_HOST:3306/$DB_NAME dbmate -e dburl $1
