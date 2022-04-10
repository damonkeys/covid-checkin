#!/bin/bash         
#
# !!!!!!!!!!!!!!!!!!!!! RUN FROM PROJECT-DIRECTORE (service-gateway) !!!!!!!!!!!!!!!!!!!!!!!!!!!!
# >> tools/deployServiceGateway.sh
#
#
# Commandline arguments
# 
# 1) will be commited to start script
#  
#
# Variables
if [ -z "$1" ]
then
    echo -e "\nERROR: Missing User for starting... eg. 'tools/deployServiceGateway.sh sho'\n\n"
    exit
fi

sudo_user=$1 #bbu or sho

server="dev.checkin.chckr.de"
server_path="/opt/ch3ck1n/"
server_user="[your-server-user]"
server_ssh=$server_user@$server
server_sudossh=$sudo_user@$server
dist_directory="dist"
current_path=$(pwd)
tools_path="../tools"

# Microserver-variables
app_name="service-gateway"
app_archive=$app_name"_SNAPSHOT_linux_amd64.tar.gz"
app_checksum=$app_name"_SNAPSHOT_checksum.txt"


echo -e "\nStarting goreleaser for building binary"
echo -e "==========================================================================\n"
#cd ../go/$1
goreleaser --snapshot --skip-validate --rm-dist
mv $dist_directory/ch3ck1n_SNAPSHOT_linux_amd64.tar.gz $dist_directory/$app_archive
mv $dist_directory/ch3ck1n_SNAPSHOT_checksums.txt $dist_directory/$app_checksum

echo -e "\nDeploy $app_directory to Server..."
echo -e "==========================================================================\n"

# Kill screen-session / Stopping server
$tools_path/stopServer.sh $app_name

# After build transfer binary to server
ssh $server_ssh "rm -Rf $server_path/$app_name"
ssh $server_ssh "mkdir $server_path/$app_name"
scp $dist_directory/$app_archive $server_ssh:$server_path/$app_name/

# Decompress..
ssh $server_ssh "tar -xvzf $server_path/$app_name/$app_archive -C $server_path/$app_name/"
ssh $server_ssh "rm $server_path/$app_name/$app_archive"

# SUDO - give access to start server with port 443
ssh -t $server_sudossh "sudo setcap CAP_NET_BIND_SERVICE=+ep $server_path/$app_name/$app_name"

# Starting server...
$tools_path/startServer.sh $app_name $2 $3 $4 $5
