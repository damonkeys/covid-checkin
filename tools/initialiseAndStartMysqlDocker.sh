#!/bin/bash
docker run -p 3306:3306 --name ch3ck1n -v $HOME/mariadb/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=REPLACE -d mariadb:latest

