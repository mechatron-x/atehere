#!/bin/bash

rm -rf logs/*

docker image prune -f
docker volume prune -f
docker network prune -f
