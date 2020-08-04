#!/bin/bash         
#
# Commandline arguments
# 
# 1) Application name       - eg. auth-server, apn-server....
#  
if [ -z "$1" ]
then
    echo -e "\nERROR: Missing Applicationname for stopping... eg. './stopServer.sh auth-server'\n\n"
    exit
fi
# Variables
server="dev.checkin.chckr.de"
server_path="/opt/ch3ck1n/"
server_user="user"
server_ssh=$server_user@$server

# Microserver-variables
app_name=$1
param2=$2

if [ -n $param2 ] && [[ $param2 = "local" ]]
    then
        # Kill screen-session
        echo -e "\nStopping service... $app_name LOCAL!"
        echo -e "==========================================================================\n"
        screen -ls | grep $app_name | cut -d. -f1 | xargs | xargs kill
        ps -A | grep $app_name | cut -d t -f1 | xargs | xargs kill
    else
        # Kill screen-session
        echo -e "\nStopping service... $app_name"
        echo -e "==========================================================================\n"
        ssh $server_ssh "screen -ls | grep $app_name | cut -d. -f1 | xargs | xargs kill"
fi
