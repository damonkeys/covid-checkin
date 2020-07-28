#!/bin/bash

docker run -p 3306:3306 --name bongodb -v $HOME/mariadb/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=monkeycashrockz2018 -d mariadb:latest
