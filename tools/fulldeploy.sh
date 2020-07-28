#!/bin/bash         

if [ -z "$1" ]
then
    echo -e "\nERROR: Missing User for starting... eg. './fulldeploy.sh sho'\n\n"
    exit
fi

echo -e "\nDeploying all monkeycash.io services"
echo -e "==========================================================================\n"

echo -e "\nSimple-Web-Server"
echo -e "==========================================================================\n"
cd ../react
yarn build
cd ../tools
./deploy.sh monkeywebserver

echo -e "\nbongo - auth-server"
echo -e "==========================================================================\n"
./deploy.sh bongo-auth

echo -e "\nkoko - qr-server"
echo -e "==========================================================================\n"
./deploy.sh koko-qr

echo -e "\nkingkong - pos-server"
echo -e "==========================================================================\n"
./deploy.sh kingkong

echo -e "\nbubbles - messaging server"
echo -e "==========================================================================\n"
./deploy.sh bubbles

### ALBERT is the last service to deploy. It has its own deploy-script!
echo -e "\nalbert - Proxy-Server"
echo -e "==========================================================================\n"
cd ../go/albert-proxy
tools/deployAlbert.sh $@
### ALBERT is the last service to deploy. It has its own deploy-script!
