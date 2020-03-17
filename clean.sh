#! /bin/bash

#docker kill $(docker ps -q)
#docker rm $(docker ps -a -q)

#if [ "$1" == "all" ]
#then 
    rm -rf ./crypto-config
    rm -rf ./network-config/*
#fi

# docker volume prune -f
# docker network prune -f

echo Done
