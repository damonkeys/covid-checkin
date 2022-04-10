#!/bin/bash

./buildEmbeddedFileServerVar.sh linux amd64
server=[your-user]@[your-host]
chckrDir=/var/opt/chckr
appName=$(basename $PWD)
gzip --best --force $(appName)
ssh -i ~/.ssh/id_rsa_monkey $(server) "killall $(appName)"
scp -Ci ~/.ssh/id_rsa_monkey landingpage.gz $(server):$(chckrDir)
ssh -i ~/.ssh/id_rsa_monkey $(server) "gzip -d --force $(chckrDir)/$(appName).gz"
ssh -i ~/.ssh/id_rsa_monkey $(server) "nohup $(chckrDir)/$(appName) > $(chckrDir)/$(appName).log 2> $(chckrDir)/$(appName)-error.log < /dev/null &"
