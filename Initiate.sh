#! /bin/bash
export FABRIC_CFG_PATH=${PWD}/config
echo -e "========>Cleanup the earlier runs====================>\n"
# ./clean.sh all

echo 
cd config

cryptogen generate --config=./crypto-config.yaml

mv -i ./crypto-config ../

configtxgen -outputBlock ../network-config/genesis.block -profile AirlineOrdererGenesis -channelID ordererchannel

configtxgen -outputCreateChannelTx ../network-config/airline-channel.tx -profile AirlineChannel -channelID airlinechannel

configtxgen -outputAnchorPeersUpdate ../network-config/AcmeAnchor.tx -channelID airlinechannel -asOrg AcmeMSP -profile AirlineChannel

# cd ..

# echo -e "===============>Launching containers================>\n"

# ./launch.sh

