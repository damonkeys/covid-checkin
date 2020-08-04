#!/bin/sh
# PATH TO YOUR HOSTS FILE
ETC_HOSTS=/etc/hosts
# DEFAULT IP FOR HOSTNAME
IP="127.0.0.1"
# Hostname to add/remove.
HOSTNAME=dev.checkin.chckr.de

function removehost() {
    if [ -n "$(grep $HOSTNAME /etc/hosts)" ]
    then
        echo "$HOSTNAME Found in your $ETC_HOSTS, Removing now...";
        sudo sed -i".bak" "/$HOSTNAME/d" $ETC_HOSTS
    else
        echo "$HOSTNAME was not found in your $ETC_HOSTS";
    fi
}

function addhost() {
    HOSTS_LINE="$IP\t$HOSTNAME"
    if [ -n "$(grep $HOSTNAME /etc/hosts)" ]
        then
            echo "$HOSTNAME already exists : $(grep $HOSTNAME $ETC_HOSTS)"
        else
            echo "Adding $HOSTNAME to your $ETC_HOSTS";
            sudo -- sh -c -e "echo '$HOSTS_LINE' >> /etc/hosts";

            if [ -n "$(grep $HOSTNAME /etc/hosts)" ]
                then
                    echo "$HOSTNAME was added succesfully \n $(grep $HOSTNAME /etc/hosts)";
                else
                    echo "Failed to Add $HOSTNAME, Try again!";
            fi
    fi
}


if [ -z "$1" ]
then
    echo -e "\nERROR: Missing command for starting... eg. './hosts.sh add'\n\n"
    exit
elif [ $1 == 'add' ]
then
    addhost
elif [ $1 == 'remove' ]
then
    removehost
else
    echo 'Unknown command. Try add or remove'
fi