#!/bin/bash         

echo -e "\nStopping all ch3ck1n services"
echo -e "==========================================================================\n"

echo -e "\nch3ck1nweb"
echo -e "==========================================================================\n"
./stopServer.sh ch3ck1nweb

echo -e "\nauth-server"
echo -e "==========================================================================\n"
./stopServer.sh auth

echo -e "\nbiz-server"
echo -e "==========================================================================\n"
./stopServer.sh biz

echo -e "\nadmin-server"
echo -e "==========================================================================\n"
./stopServer.sh admin

echo -e "\nservice-gateway"
echo -e "==========================================================================\n"
./stopServer.sh service-gateway
