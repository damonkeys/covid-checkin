#!/bin/bash

#needed: ./run.sh up or ./run.sh migrate
# generally speaking: ./run.sh <databasename> <dbmate_command>
dburl=mysql://auth_user:@127.0.0.1:3306/$1 dbmate -e dburl $2
