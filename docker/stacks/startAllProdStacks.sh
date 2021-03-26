#! /bin/bash
cd $(dirname "$0")

./chckr/startProdStack.sh
./homepage/startProdStack.sh
./landing/startStack.sh
./reverse-proxy/startProdStack.sh
