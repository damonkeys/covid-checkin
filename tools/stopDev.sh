#!/bin/sh
# Script for stopping all servers locally!

./hosts.sh remove

./stopServer.sh service-gateway local
./stopServer.sh auth local
./stopServer.sh biz local
./stopServer.sh admin local
./stopServer.sh ch3ck1nweb local
./stopServer.sh pixi local
