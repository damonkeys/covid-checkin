#!/bin/bash         
#
# Commandline arguments
# 
# 1) Application name       - eg. auth-server, apn-server....
#  
if [ -z "$1" ]
then
    echo -e "\nERROR: Missing Applicationname for starting... eg. './startServer.sh auth'\n\n"
    exit
fi

app_name=$1
param2=$2

if [ -n $param2 ] && [[ $param2 = "local" ]]
    then
        # local environment - development
        echo -e "\nStarting service... $1 LOCAL!"
        echo -e "==========================================================================\n"
        screen -S $app_name -dmS bash -c "cd ../$app_name/; go build; ./startServer.sh; exec bash"
    else
        # Variables
        server="dev.checkin.chckr.de"
        server_path="/opt/ch3ck1n/"
        server_user="user"
        server_ssh=$server_user@$server

        # Starting server..."
        echo -e "\nStarting service... $app_name"
        echo -e "==========================================================================\n"
        ssh $server_ssh "screen -S $app_name -dm bash -c 'cd $server_path/$app_name/; ./startServer.sh; exec sh'"
fi
