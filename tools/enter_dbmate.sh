#! /bin/bash
docker run --name chckr_dbmate --rm -it -v "$(pwd)/db:/db" --network=chckr_chckr_default chckr/dbmate
