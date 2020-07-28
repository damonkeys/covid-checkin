#!/bin/sh
# Script for stopping all servers locally!

./hosts.sh remove

./stopServer.sh albert-proxy local
./stopServer.sh bongo-auth local
./stopServer.sh koko-qr local
./stopServer.sh react local
./stopServer.sh kingkong local
./stopServer.sh bubbles local
