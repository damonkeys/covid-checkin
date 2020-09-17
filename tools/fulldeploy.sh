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

echo -e "\nbiz-server"
echo -e "==========================================================================\n"
./deploy.sh biz

echo -e "\npixi-server"
echo -e "==========================================================================\n"
./deploy.sh pixi

echo -e "\nadmin-server"
echo -e "==========================================================================\n"
./deploy.sh admin

### Service-Gateway is the last service to deploy. It has its own deploy-script!
echo -e "\Service-Gateway"
echo -e "==========================================================================\n"
cd ../service-gateway
tools/deployServiceGateway.sh $@
### Service-Gateway is the last service to deploy. It has its own deploy-script!
