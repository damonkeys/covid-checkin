# Docker
All microservices are dockerized. You can start the whole development-environment with docker-compose to bring up all
servers, app and all other needed servers like MariaDB and KONG.

## First time
Starting the docker environment for the first time you have to build all Docker Images before.

In the tools-folder you find a script for starting the docker-environment called `startDocker.sh`. Call this command at the first time to build the whole environment and start all containers.

```
cd tools
./startDocker.sh dev build
```
For the dev-stage you are not able to reach the react-app becaus kong is configured to call the local running app. You have to start it manually via
```
cd client-app
yarn start
```
It will be started locally because of developing without deploying the react-app. Now you are ready to call the app via https://localhost.

To stop all docker containers call the `stopDocker.sh` script.

```
cd tools
./stopDocker.sh dev
```

### Stage prod
Stage prod works the same as dev but it starts ch3ck1nweb-server with static build of the react-app. You are ready to reach chckr via the prod-URL defined in `docker/env/.env.prod`.

```
cd tools
./startDocker prod build
```

### Other stages
You can define freely any stage you need. The env-files and Kong-config is grouped by folders named like the given stage.


### Database
For the first time starting the docker-composer envrionment there will be an instance of MariaDB created but without any tables. Please run the following dbmate-scripts to create all necessary tables and users:

```
cd authx/dbmate
./run.sh up

cd ../../biz/dbmate
./run.sh up

cd ../../checkins/dbmate
./run.sh up
```

## Second time ;-)
For the second time you don't need tob build or initialize something. Enter the tools folder and start the docker-environment:
```
cd tools
./startDocker dev

cd ../client-app
yarn start
```

Now you are ready for developing!


# Scripts

## Build-Scripts
There is a script for building an image in this docker folder:

```
cd docker
./buildDockerImage.sh <servicename>
```
To rebuild a single golang-image after changing the code you need this command e.g.
```
cd docker
./buildDockerImage.sh biz
```

To build all containers there is a script for multiple call the buildDockerImage.sh script. _This script is needed to add new servers if we add one to our environment!_
```
cd docker
./buildAll.sh
```

## Start-Scripts
For starting the docker-containers in tools folder you find the `startDocker.sh` script.

```
cd tools
./startDocker.sh <stage> <build>
```

The two parameters are for controlling the starting.

1. The first parameter stage expect dev or prod. With this parameter you control some environment-variables for the docker-container. You find this environment-variables in the folder `docker/env`. For both stages there are one env-file:
    ```
    .env.dev
    .env.prod
    ```
    Here you find the Domainname for this stage.

    In dependency of the stage kong gets the specific route-config. You find both configs in folder `docker/kong`. The `kong/dev/kong.yml` has to be changed locally for your development. This is described later.

2. The second parameter is optional. If you set it to build all images will be (re)build. Sometimes it is necessary to rebuild the whole environment.

# Developing with docker
For developing it is very slow to rebuild the docker iamge for testing your changes on the code. If you do some changes to a server start it locally like the old way with the `./startServer.sh` script in your server-folder, e.g.
```
cd biz
./startServer.sh
```
At this moment kong doesn't knows your local server and will call the docker-container via the route config. Before you start the whole docker-environment you have to adjust the kong-route-file. You find it in `docker/kong/dev/kong.yml`.

This file is prepared for development-routes. For example if you want to use the local running biz server you have to tell it kong. change the kong-route section from:
```
- name: biz
  url: http://biz:4000
  routes:
    - name: biz-route
      paths:
        - /biz
```
to this:
```
- name: biz
  url: http://host.docker.internal:4000
  routes:
    - name: biz-route
      paths:
        - /biz
```
Now kong will route all biz-calls to your locally running server.

After this adjustments you start the docker environment:
```
cd tools
./startDocker dev
```
All your biz-server-changes can be tested locally. Then you have to rebuild and restart the server locally:
```
cd <servername>
go build
./startServer.sh
```

## Logging
The `./startDocker.sh`script starts the docker-composer container in detached mode. You won't see any logs. To get the
logs of a service try this command:
```
docker-compose -f ./docker/docker-compose.yml -p chckr logs -f biz
```
Or use the simple script `dockerLog.sh <servicename>` from tools folder.

## changing Kong-route-config: kong-yml
We use kong database-less so it uses a simple config-file for all route-definition. You will find it in `docker/kong/<stage>`. After adjusting the routes in kong.yml you have to tell kong to reload the configuration. You will do it with curl:
```
curl --data-urlencode "config@../docker/kong/dev/kong.yml" -X POST http://localhost:8001/config
```
Or simple with the reload-script in tools folder:
```
cd tools
./relaodKongConfig.sh
```
