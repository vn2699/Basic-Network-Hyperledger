export CC_CONSTRUCTOR1='{"function":"initLedger","Args":[""]}'
export CC_CONSTRUCTOR2='{"Args":[]}'
export CC_NAME="gocc"
export CC_PATH="github.com/chaincode/"
export CC_VERSION="1.0"
export CC_CHANNEL_ID="airlinechannel"

OPERATION=$1

echo "CC OPERATION : $OPERATION"

case $OPERATION in
    "install" )
        peer chaincode install -n $CC_NAME -v $CC_VERSION -p $CC_PATH
        peer chaincode install -n gocc -v 1.0 -p "github.com/chaincode/"
        peer chaincode list --installed
        ;;
    "instantiate" )
        peer chaincode instantiate -n $CC_NAME -v $CC_VERSION -C $CC_CHANNEL_ID -c $CC_CONSTRUCTOR2
        peer chaincode list --instantiated
        ;;
    "invoke" )
        peer chaincode invoke -n $CC_NAME -C $CC_CHANNEL_ID -c $CC_CONSTRUCTOR1
        ;;
    "query" )
        peer chaincode query  -n $CC_NAME -C $CC_CHANNEL_ID -c '{"function":"queryAirline","Args":["1"]}'
        ;;
    * )
        echo "Provide the right operations"
        ;;
esac

