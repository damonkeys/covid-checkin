#!/bin/bash         

echo -e "\nStarting all ch3ck1n services"
echo -e "==========================================================================\n"

echo -e "\nproxy-server"
echo -e "==========================================================================\n"
../proxy/tools/startServer.sh proxy

echo -e "\nch3ck1nweb"
echo -e "==========================================================================\n"
./startServer.sh ch3ck1nweb 

echo -e "\nauth-server"
echo -e "==========================================================================\n"
./startServer.sh auth
