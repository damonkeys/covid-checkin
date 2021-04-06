#! /bin/bash
docker run --name chckr_dbmate --rm -it -v "$(pwd)/db:/db" --network=chckr_default ${{ secrets.REGISTRY_SERVER }}/chckr/dbmate
