#!/bin/bash
docker-compose -f ./docker-compose-all.yaml down

docker-compose -f ./docker-compose-cli.yaml down

echo "Checking if all docker Containers are down"

docker ps -a

echo "Done"
