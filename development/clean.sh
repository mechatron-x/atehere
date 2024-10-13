#!/bin/bash

docker image prune -f
docker volume prune -f
docker network prune -f
