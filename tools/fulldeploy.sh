#!/bin/bash         

if [ -z "$1" ]
then
    echo -e "\nERROR: Missing User for starting... eg. './fulldeploy.sh sho'\n\n"
    exit
fi

echo -e "\nDeploying all ch3ck1n services"
echo -e "==========================================================================\n"

echo -e "\ch3ck1nweb-Web-Server"
echo -e "==========================================================================\n"
cd ../client-app
yarn build
cd ../tools
./deploy.sh ch3ck1nweb

echo -e "\nauth-server"
echo -e "==========================================================================\n"
./deploy.sh auth

### ALBERT is the last service to deploy. It has its own deploy-script!
echo -e "\Proxy-Server"
echo -e "==========================================================================\n"
cd ../proxy
tools/deployProxy.sh $@
### ALBERT is the last service to deploy. It has its own deploy-script!
