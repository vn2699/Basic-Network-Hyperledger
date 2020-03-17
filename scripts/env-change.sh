#! /bin/bash

if [ -z $1 ]
then
    echo Enter the name of the organzation on behalf to run the scripts, either acme or budget
    exit 1
else
    ORG_NAME2=${1,,}
fi
# export CORE_PEER_ID=peer0.$ORG_NAME2.com

export ORG_NAME2=acme
export CORE_PEER_CHAINCODELISTENADDRESS=peer0.$ORG_NAME2.com:7052
export CORE_PEER_ADDRESS=peer0.$ORG_NAME2.com:7051
export CORE_PEER_TLS_ENABLED=false
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG_NAME2.com/users/Admin@$ORG_NAME2.com/msp
export ORG_NAME1=${ORG_NAME2^}
# echo $ORG_NAME1
export CORE_PEER_LOCALMSPID="${ORG_NAME1}MSP"
# export ORG_NAME=peer0.$ORG_NAME2.com    
# export ORDERER_ADDRESS=orderer.acme.com:7050
export ORG_NAME2=budget
export CORE_PEER_CHAINCODELISTENADDRESS=peer0.$ORG_NAME2.com:8052
export CORE_PEER_ADDRESS=peer0.$ORG_NAME2.com:8051
export CORE_PEER_TLS_ENABLED=false
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/$ORG_NAME2.com/users/Admin@$ORG_NAME2.com/msp
export ORG_NAME1=${ORG_NAME2^}
# echo $ORG_NAME1
export CORE_PEER_LOCALMSPID="${ORG_NAME1}MSP"
echo -e "==========>ENV VARIABLES<==============\n"
env | grep CORE_
