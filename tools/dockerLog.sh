#! /bin/bash
STAGE=dev docker-compose -f ../docker/docker-compose.yml -p chckr logs -f $1
