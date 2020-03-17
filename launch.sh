#! /bin/bash

docker-compose -f ./docker-compose-all.yaml up -d

# docker-compose -f ./docker-compose-cli.yaml up -d

echo "Network started"
echo "========================================"
echo "Checking the containers"
docker ps -a
echo "All containers running successfully..."
echo "========================================"
