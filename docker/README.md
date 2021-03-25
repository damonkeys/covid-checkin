# Docker
All microservices are dockerized. You can start the chckr-environment with docker stacks to bring up the whole swarm. 
We have different stacks for chckr-services, landing-pages, website and a global reverse proxy. All stacks are defined
by seperate docker-folder for compose-files in folder docker/stacks/.

## First time
Starting the docker environment for the first time you have to build all Docker images before. In production this is done
by github actions or locally use the script "buildAll.sh" in tools-folder. But to start the docker
swarm you need to init it with following docker command:

```
docker swarm init
```

Before starting the stacks the first time you need to start some scripts. Create needed networks using the script in
docker/onetime/. Use this for local development:

```
docker/onetime/createNetworks.sh
docker/onetime/addDevHosts.sh
```

or on remote docker server (e.g. for production)
```
DOCKER_HOST="ssh://user@dockerhost" docker/onetime/createNetworks.sh
```

We need some config-files. In development-environment they are located relativ to the stack-folder. There is nothing to
do in local development-environment. On docker-server in production you need to copy the config-files to the expected
locations. Execute the script copyProdConfigFiles.sh located in onetime folder:

```
SSH_HOST="user@dockerhost" docker/onetime/copyProdConfigFiles.sh
```

### Stage dev
For the local-development-environemnt you are not able to reach the react-app because albert is configured to call the
locally running app. You have to start it manually via
```
cd client-app
yarn start
```
It will be started locally because of developing without deploying the react-app. Now you are ready to call the app via
https://dev.checkin.chckr.de. __Don't forget to call the ./host.sh add script.__

To start all other services you need to deploy all needed stacks:

```
docker stack deploy -c docker/stacks/reverse-proxy/docker-compose.dev.yml --with-registry-auth proxy
docker stack deploy -c docker/stacks/homepage/docker-compose.dev.yml --with-registry-auth www
docker stack deploy -c docker/stacks/landing/docker-compose.yml --with-registry-auth landing
cd docker/stacks/chckr
./deployDevStack.sh
```


### Stage prod
Stage prod works the same as dev but it starts ch3ck1nweb-server with static build of the react-app.

```
DOCKER_HOST="ssh://user@dockerhost" docker stack deploy -c docker/stacks/reverse-proxy/docker-compose.prod.yml --with-registry-auth proxy
DOCKER_HOST="ssh://user@dockerhost" docker stack deploy -c docker/stacks/homepage/docker-compose.prod.yml --with-registry-auth www
DOCKER_HOST="ssh://user@dockerhost" docker stack deploy -c docker/stacks/landing/docker-compose.yml --with-registry-auth landing
cd docker/stacks/chckr
DOCKER_HOST="ssh://user@dockerhost" ./deployProdStack.sh
```


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


# Scripts

## Build-Scripts
There is a script for building an image in the tools folder:

```
cd tools
./buildDockerImage.sh <servicename>
```
To rebuild a single golang-image after changing the code you need this command e.g.
```
cd docker
./buildDockerImage.sh biz
```

To build all containers there is a script for multiple call the buildDockerImage.sh script. _This script is needed to
add new servers if we add one to our environment!_ The script builds the dbmate-image with all migrations too. __No docker images will be pushed to the docker registry!!!__
```
cd tools
./buildAll.sh
```

# Developing with docker
For developing it is very slow to rebuild the docker iamge for testing your changes on the code. If you do some changes to a server start it locally like the old way with the `./startServer.sh` script in your server-folder, e.g.
```
cd biz
./startServer.sh
```
At this moment Albert doesn't know your local server and will call the container via the route config. Before you start the whole docker-environment you have to adjust the Albert-route-file. You find it in `docker/stacks/chckr/albert/dev/route.json`.

For example if you want to use the local running biz server you have to tell it Albert. Change the Albert-route from:
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

After this adjustments you have to restart the docker-dev-stack.

All your biz-server-changes can be tested locally. Then you have to rebuild and restart the server locally:
```
cd <servername>
go build
./startServer.sh
```

## Logging
The docker swarm containers running in detached mode. You won't see any logs. To get the
logs of a service try this command:
```
docker service logs -f chckr_biz
```
Or use the simple script `dockerLog.sh <servicename>` from tools folder.
```
./dockerLog.sh biz
```
