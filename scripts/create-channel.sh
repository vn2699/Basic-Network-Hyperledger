#! /bin/bash

peer channel create -o orderer.acme.com:7050 -c airlinechannel -f ../network-config/airline-channel.tx
peer channel join -b airlinechannel.block
