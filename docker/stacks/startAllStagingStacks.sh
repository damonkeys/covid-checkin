#! /bin/bash
cd $(dirname "$0")

./chckr/startStagingStack.sh
./homepage/startStagingStack.sh
./landing/startStack.sh
./reverse-proxy/startStagingStack.sh
