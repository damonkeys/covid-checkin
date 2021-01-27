#!/usr/bin/env bash
docker build -t landing-$(basename $PWD) .
