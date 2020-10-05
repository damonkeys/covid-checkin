#!/bin/sh
# Script for starting all servers locally!

./hosts.sh add

./startServer.sh service-gateway local
./startServer.sh authx local
./startServer.sh biz local
./startServer.sh admin local
./startServer.sh pixi local
./startServer.sh checkins local


echo -e "\nStarting service... react/ client-app LOCAL!"
echo -e "==========================================================================\n"
screen -S ch3ck1nweb -dmS bash -c "cd ../client-app/; yarn start; exec bash"
