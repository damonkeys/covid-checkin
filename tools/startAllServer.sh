#!/bin/bash         

echo -e "\nStarting all monkeycash.io services"
echo -e "==========================================================================\n"

echo -e "\nalbert - proxy-server"
echo -e "==========================================================================\n"
../go/albert-proxy/tools/startServer.sh albert-proxy

echo -e "\nmonkeycashweberver"
echo -e "==========================================================================\n"
./startServer.sh monkeywebserver 

echo -e "\nbongo - auth-server"
echo -e "==========================================================================\n"
./startServer.sh bongo-auth

echo -e "\nkoko - qr-server"
echo -e "==========================================================================\n"
./startServer.sh koko-qr
