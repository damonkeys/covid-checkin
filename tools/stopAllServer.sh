#!/bin/bash         

echo -e "\nStopping all ch3ck1n services"
echo -e "==========================================================================\n"

echo -e "\nch3ck1nweb"
echo -e "==========================================================================\n"
./stopServer.sh ch3ck1nweb

echo -e "\nauth-server"
echo -e "==========================================================================\n"
./stopServer.sh auth

echo -e "\nproxy server"
echo -e "==========================================================================\n"
./stopServer.sh proxy
