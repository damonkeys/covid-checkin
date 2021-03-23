#! /bin/bash
docker network create -d overlay chckr_default --scope swarm
docker network create -d overlay landing_default --scope swarm
docker network create -d overlay homepage_default --scope swarm
