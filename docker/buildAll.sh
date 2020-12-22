#! /bin/bash
./buildDockerImage.sh admin
./buildDockerImage.sh authx
./buildDockerImage.sh biz
./buildDockerImage.sh checkins
./buildDockerImage.sh pixi
./buildDockerImage.sh service-gateway

# Cleanup ch3ck1nweb-static-folder
cd ../ch3ck1nweb
rm -Rf static/
mkdir static

# Build yarn app and copy all static files
cd ../client-app/
yarn build
cp -R build/* ../ch3ck1nweb/static/

cd ../docker
./buildDockerImage.sh ch3ck1nweb

# dbmate container
cd ../dbmate
docker build -t chckr/dbmate -t ${{ secrets.REGISTRY_SERVER }}/chckr/dbmate .
