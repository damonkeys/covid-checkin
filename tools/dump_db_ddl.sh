#!/bin/bash

# dumps the given db structure without data and statistics for 127.0.0.1 to $1_schema.sql.
# usage: $1 dbname $2 dbuser
mysqldump --databases $1 -h 127.0.0.1 -u $2 -p --no-create-db --no-data --column-statistics=0 > $1_schema.sql
