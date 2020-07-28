#!/bin/bash         
#
# Commandline arguments
# 
# 1) Application name       - eg. auth-server, apn-server....
#  
if [ -z "$1" ]
then
    echo -e "\nERROR: Missing Applicationname for deploying... eg. './deploy.sh auth-server'\n\n"
    exit
fi
# Variables
server="checkin.chckr.de"
server_path="/opt/monkeycash/"
server_user="pmd"
server_ssh=$server_user@$server
dist_directory="dist"
current_path=$(pwd)

# Microserver-variables
app_name=$1
app_archive="$1_SNAPSHOT_linux_amd64.tar.gz"
app_checksum="$1_SNAPSHOT_checksum.txt"


echo -e "\nStarting goreleaser for building binary"
echo -e "==========================================================================\n"
cd ../go/$1
goreleaser --snapshot --skip-validate --rm-dist
mv $dist_directory/monkeycash_SNAPSHOT_linux_amd64.tar.gz $dist_directory/$app_archive
mv $dist_directory/monkeycash_SNAPSHOT_checksums.txt $dist_directory/$app_checksum

echo -e "\nDeploy $app_directory to Server..."
echo -e "==========================================================================\n"

# Kill screen-session / Stopping server
$current_path/stopServer.sh $app_name

# After build transfer binary to server
ssh $server_ssh "rm -Rf $server_path/$app_name"
ssh $server_ssh "mkdir $server_path/$app_name"
scp $dist_directory/$app_archive $server_ssh:$server_path/$app_name/

# Decompress..
ssh $server_ssh "tar -xvzf $server_path/$app_name/$app_archive -C $server_path/$app_name/"
ssh $server_ssh "rm $server_path/$app_name/$app_archive"

# Starting server...
$current_path/startServer.sh $app_name $2
