#! /bin/bash

container_id=$(docker ps -aqf "name=dbmate_dbmate")
echo 'Container-ID for dbmate-container: ' $container_id
docker exec -it $container_id bash
