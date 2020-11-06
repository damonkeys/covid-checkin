#! /bin/bash
./buildDockerImage.sh admin
./buildDockerImage.sh authx
./buildDockerImage.sh biz
./buildDockerImage.sh checkins
./buildDockerImage.sh pixi
cd ../client-app/; yarn build; cd ../docker; ./buildDockerImage.sh ch3ck1nweb
