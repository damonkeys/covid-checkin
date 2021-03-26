#! /bin/bash
cd $(dirname "$0")

./chckr/startDevStack.sh
./homepage/startDevStack.sh
./landing/startStack.sh
./reverse-proxy/startDevStack.sh
