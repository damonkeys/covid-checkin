#!/bin/sh
# PATH TO YOUR HOSTS FILE
ETC_HOSTS=/etc/hosts
# DEFAULT IP FOR HOSTNAME
IP="127.0.0.1"
# Hostname to add/remove.
HOSTNAME=chckr.local


if [ -n "$(grep $HOSTNAME /etc/hosts)" ]
    then
        echo "$HOSTNAME already exists : $(grep $HOSTNAME $ETC_HOSTS)"
    else
        echo "Adding $HOSTNAME to your $ETC_HOSTS";
        sudo -- sh -c -e "echo '$IP\t$HOSTNAME' >> /etc/hosts";
        sudo -- sh -c -e "echo '$IP\twww.$HOSTNAME' >> /etc/hosts";
        sudo -- sh -c -e "echo '$IP\tlanding.$HOSTNAME' >> /etc/hosts";
        sudo -- sh -c -e "echo '$IP\tcheckin.$HOSTNAME' >> /etc/hosts";

        if [ -n "$(grep $HOSTNAME /etc/hosts)" ]
            then
                echo "$HOSTNAME was added succesfully \n $(grep $HOSTNAME /etc/hosts)";
            else
                echo "Failed to Add $HOSTNAME, Try again!";
        fi
fi
