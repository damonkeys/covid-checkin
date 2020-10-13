#!/bin/bash

# https://hub.docker.com/_/mariadb
docker run -p 3306:3306 --name ch3ck1n -v $HOME/mariadb/data_ch3ckin:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=REPLACE \
    -e MYSQL_DATABASE=ch3ck1n \
    -e MYSQL_USER=ch3ck1n_user \
    -e MYSQL_PASSWORD===> \
    -d mariadb:latest
