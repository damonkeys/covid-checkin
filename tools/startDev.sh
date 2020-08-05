#!/bin/sh
# Script for starting all servers locally!

./hosts.sh add

./startServer.sh service-gateway local
./startServer.sh auth local

screen -S react -dmS bash -c "cd ../client-app/; yarn start; exec bash"
