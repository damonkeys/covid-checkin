#!/bin/sh
# Script for starting all servers locally!

./hosts.sh add

./startServer.sh albert-proxy local
./startServer.sh bongo-auth local
./startServer.sh koko-qr local
./startServer.sh kingkong local
./startServer.sh bubbles local

screen -S react -dmS bash -c "cd ../react/; yarn start; exec bash"
