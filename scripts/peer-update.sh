#! /bin/bash

ORG_NAME=$1

PEER_FABRIC_CFG_PATH=/opt/gopath/src/github.com/hyperledger/fabric/peer

./env-change.sh acme

peer channel signconfigtx -f ./network-config/AcmeAnchor.tx

./env-change.sh budget

peer channel signconfigtx -f ./network-config/AcmeAnchor.tx

# configtxgen -outputAnchorPeersUpdate /opt/gopath/src/github.com/hyperledger/fabric/peer/network-config -channelID airlinechannel -asOrg $ORG_NAME -profile AirlineChannel

peer channel update -f /opt/gopath/src/github.com/hyperledger/fabric/peer/network-config/AcmeAnchor.tx -c airlinechannel -o orderer0.acme.com:7050