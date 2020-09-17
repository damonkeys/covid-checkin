#!/bin/bash         

echo -e "\nStarting all ch3ck1n services"
echo -e "==========================================================================\n"

echo -e "\nservice-gateway"
echo -e "==========================================================================\n"
../service-gateway/tools/startServer.sh service-gateway

echo -e "\nch3ck1nweb"
echo -e "==========================================================================\n"
./startServer.sh ch3ck1nweb 

echo -e "\nbiz"
echo -e "==========================================================================\n"
./startServer.sh biz

echo -e "\npixi"
echo -e "==========================================================================\n"
./startServer.sh pixi

echo -e "\nadmin"
echo -e "==========================================================================\n"
./startServer.sh admin 

echo -e "\nauth-server"
echo -e "==========================================================================\n"
./startServer.sh auth
