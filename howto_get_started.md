# How to get started with chckr as a devloper

I little how to help you wire up your dev environment

## Steps

### Preparation

You should have a decent package manager running through which you should install/ have installed:

 * go/golang
 * npm
 * git
 * docker <- might not work through package manager. Do the right thing for your operating system
 * vscode

### Initialisation - the docker way
* clone the repo
* run the command `docker swarm init`
* in the tools folder run the script `./startStack.sh dev build`
* in the tools folder run the script `enter_dbmate.sh`
* change in the dbmate-docker container to the directories _db-chckr/ and db-checkins_ and run the script `./run.sh up` to execute all migrations.

For further informations about the docker environment look at the README.md at the docker folder!

### Initialisation - the old way

* clone the repo
* in the tools folder run the script `./deprecated/initialiseAndStartMariadbDocker.sh`
* Edge case for folder _/checkins_: We need to initialise the seperate db first. The script `createDBAndUser.sh` in the _/checkins_ folder should do that.
* in the tools folder run the script `install_dbmate.sh`
* check the directories  _admin/ authx/ /biz /checkins_ for a folder _/dbmate_ and run the script `./run.sh [database]` with the correct database name from within that folder.
_If you don't know the correct name ask a developer/ mentor_

### First run - the docker way
* change into _/tools_ folder and start the script `./host.sh add` to add an ip-routing for checkins.monkeycash.io to your localhost.
* change into _/tools_ folder. restart all docker containers with `./stopStack.sh dev; ./startStack.sh dev`
* change into _/client-app folder to start the react-app with `yarn start`

In general your default system browser should be open itself and try to `GET` _localhost:3000_ <- this is wrong unfortunately. use the https-albert-route and start the app with https://dev.checkin.chckr.de

Your browser should render the checkin service by now.

For further informations about the docker environment look at the README.md at the docker folder!

*Congratulation! You made it. Welcome on board developer*

### First run - the old way

* change into _/tools folder and run `./startDev.sh`
* check with `screen -ls` and optional `screen -r [screename]` wether you have "the services" running

In general your default system browser should be open itself and try to `GET` _localhost:3000_ <- this is wrong unfortunately:

The ./startDev.sh script fiddles with local _hostnames_ so that you can/ must use https://dev.checkin.chckr.de as a starting point into the app.

Your browser should render the checkin service by now.

__deprecated-scripts:__ Because of dockerizing process of all servers some scripts in the tools-folder moved to the
deprecated folder. In future maybe all scripts of the the old way methode will be moved or deleted.

*Congratulation! You made it. Welcome on board developer*

![](https://media3.giphy.com/media/bznNJlqAi4pBC/giphy.gif)
