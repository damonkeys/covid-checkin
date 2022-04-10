#! /bin/bash

if [ -z "$SSH_HOST" ]
then
    echo -e "\nERROR: please execute this script with SSH_HOST set to target-host."
    echo -e ""
    echo -e 'example: SSH_HOST="[your-user]@[your-dockerhost]" ./copyProdConfigFiles.sh'
    exit
fi

ssh $SSH_HOST -t "mkdir -p ~/chckr-config-copy"

# reverse-proxy
ssh $SSH_HOST -t "mkdir -p ~/chckr-config-copy/etc/chckr/caddy"
scp ../stacks/reverse-proxy/caddy/prod/Caddyfile $SSH_HOST:~/chckr-config-copy/etc/chckr/caddy/

# homepage
ssh $SSH_HOST -t "mkdir -p ~/chckr-config-copy/etc/chckr/php"
scp ../stacks/homepage/php/prod/homepage.ini $SSH_HOST:~/chckr-config-copy/etc/chckr/php/

# chckr
ssh $SSH_HOST -t "mkdir -p ~/chckr-config-copy/etc/chckr/albert"
scp ../stacks/chckr/albert/prod/routes.json $SSH_HOST:~/chckr-config-copy/etc/chckr/albert/

# sudo
ssh $SSH_HOST -t "sudo mkdir -p /var/log/chckr; sudo mv ~/chckr-config-copy/etc/chckr /etc/; sudo chown -R root:root /etc/chckr"
ssh $SSH_HOST -t "rm -Rf ~/chckr-config-copy"
