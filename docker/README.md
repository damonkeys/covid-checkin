# Docker
All microservices are dockerized. You can start the chckr-environment with docker stacks to bring up the whole swarm.

## First time
Starting the docker environment for the first time you have to build all Docker Images before. But to start the docker
swarm you need to init it with following docker command:

```
docker swarm init
```

In the tools-folder you find a script for starting the docker-environment called `startStack.sh`. Call this command at
the first time to build the whole environment and start all containers.

```
cd tools
./startStack.sh dev build
```
For the dev-stage you are not able to reach the react-app because kong is configured to call the local running app. You
have to start it manually via
```
cd client-app
yarn start
```
It will be started locally because of developing without deploying the react-app. Now you are ready to call the app via
https://dev.checkin.chckr.de. __Don't forget to call the ./host.sh add script.__

To stop all docker containers call the `stopStack.sh` script.

```
cd tools
./stopStack.sh
```

### Stage prod
Stage prod works the same as dev but it starts ch3ck1nweb-server with static build of the react-app. You are ready to reach chckr via the prod-URL defined in `docker/env/.env.prod`.

```
cd tools
./startStack.sh prod build
```

### Other stages
You can define freely any stage you need. The env-files and Albert-config are grouped by folders named like the given stage.


### Database
During the first start of the docker environment there will be two clean instances of MariaDB created but without any
tables. To run the dbmate-scripts for migrations you need to start another docker container called chckr/dbmate. To enter
the commandline of this container there is a script in the tools-directory:

```
cd tools
./enter_dbmate.sh
```

Then you are in the bash of the dbmate container. There are two directories: db-chckr and db-checkins. These are all
migrations for both databases. You need to enter the directories for each database and run the following script:
```
cd /db-chckr
./run.sh up

cd /db-checkins
./run.sh up
```

## Second time ;-)
During the second time you don't need to build or initialize something. Enter the tools folder and start the docker-environment:
```
cd tools
./startStack.sh dev

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

To build all containers there is a script for multiple call the buildDockerImage.sh script. _This script is needed to
add new servers if we add one to our environment!_ The script builds the dbmate-image with all migrations too.
```
cd docker
./buildAll.sh
```

## Start-Scripts
To start the containers in the tools folder you find the `startStack.sh` script.

```
cd tools
./startStack.sh <stage> <build>
```

The two parameters are for controlling the starting.

1. The first parameter stage expect dev or prod. With this parameter you control some environment-variables for the container. You find this environment-variables in the folder `docker/env`. For both stages there are one env-file:
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
At this moment Albert doesn't know your local server and will call the container via the route config. Before you start the whole docker-environment you have to adjust the Albert-route-file. You find it in `docker/albert/dev/route.json`.

For example if you want to use the local running biz server you have to tell it Albert. change the Albert-route from:
```
{
  "name": "Biz",
  "path": "/biz",
  "description":"The route to the biz server",
  "urls": [
      "http://biz:4000"
  ],
  "balancer": "roundrobin",
  "rewrite": true
}
```
to this:
```
{
  "name": "Biz",
  "path": "/biz",
  "description":"The route to the biz server",
  "urls": [
      "http://host.docker.internal:4000"
  ],
  "balancer": "roundrobin",
  "rewrite": true
}
```
Now Albert will route all biz-calls to your locally running server.

After this adjustments you start the docker environment:
```
cd tools
./startStack.sh dev
```
All your biz-server-changes can be tested locally. Then you have to rebuild and restart the server locally:
```
cd <servername>
go build
./startServer.sh
```

## Logging
The `./startStack.sh`script starts the docker swarm containers in detached mode. You won't see any logs. To get the
logs of a service try this command:
```
docker service logs -f chckr_biz
```
Or use the simple script `dockerLog.sh <servicename>` from tools folder.
```
./dockerLog.sh biz
```
