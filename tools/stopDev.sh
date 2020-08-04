#!/bin/sh
# Script for stopping all servers locally!

./hosts.sh remove

./stopServer.sh proxy local
./stopServer.sh auth local
