#!/bin/bash         

echo -e "\nStopping all ch3ck1n services"
echo -e "==========================================================================\n"

echo -e "\SimpleWebServer"
echo -e "==========================================================================\n"
./stopServer.sh simplewebserver

echo -e "\nauth-server"
echo -e "==========================================================================\n"
./stopServer.sh auth

echo -e "\nproxy server"
echo -e "==========================================================================\n"
./stopServer.sh proxy
