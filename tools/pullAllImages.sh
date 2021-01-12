#/bin/bash

echo "updates all existing images from registry"
docker images --format "{{.Repository}}:{{.Tag}}" | grep ':latest' | xargs -L1 sudo docker pull;
