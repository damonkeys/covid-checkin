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

### Initialisation

* clone the repo
* in the tools folder run the script `./initialiseAndStartMariadbDocker.sh`
* Edge case for folder _/checkins_: We need to initialise the seperate db first. The script `createDBAndUser.sh` in the _/checkins_ folder should do that.
* in the tools folder run the script `install_dbmate.sh`
* check the directories  _admin/ authx/ /biz /checkins_ for a folder _/dbmate_ and run the script `./run.sh [database]` with the correct database name from within that folder.
_If you don't know the correct name ask a developer/ mentor_

### First run

* change into _/tools folder and run `./startDev.sh`
* check with `screen -ls` and optional `screen -r [screename]` wether you have "the services" running

In general your default system browser should be open itself and try to `GET` _localhost:3000_ <- this is wrong unfortunately:

The ./startDev.sh script fiddles with local _hostnames_ so that you can/ must use https://dev.checkin.chckr.de as a starting point into the app.

Your browser should render the checkin service by now.

*Congratulation! You made it. Welcome on board developer*

![](https://media3.giphy.com/media/bznNJlqAi4pBC/giphy.gif)
