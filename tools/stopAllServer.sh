#!/bin/bash         

echo -e "\nStopping all monkeycash services"
echo -e "==========================================================================\n"

echo -e "\nMonkeyWebServer"
echo -e "==========================================================================\n"
./stopServer.sh monkeywebserver

echo -e "\nbongo - auth-server"
echo -e "==========================================================================\n"
./stopServer.sh bongo-auth

echo -e "\nkoko - qr-server"
echo -e "==========================================================================\n"
./stopServer.sh koko-qr

echo -e "\nbubbles - messagig-server"
echo -e "==========================================================================\n"
./stopServer.sh bubbles

echo -e "\nalbert - proxy server"
echo -e "==========================================================================\n"
./stopServer.sh albert-proxy
