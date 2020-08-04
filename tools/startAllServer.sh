#!/bin/bash         

echo -e "\nStarting all ch3ck1n services"
echo -e "==========================================================================\n"

echo -e "\nproxy-server"
echo -e "==========================================================================\n"
../proxy/tools/startServer.sh proxy

echo -e "\nsimplewebserver"
echo -e "==========================================================================\n"
./startServer.sh simplewebserver 

echo -e "\nauth-server"
echo -e "==========================================================================\n"
./startServer.sh auth
