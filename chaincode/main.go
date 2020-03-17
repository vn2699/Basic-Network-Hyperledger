package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Airline struct {
	Parts    string `json:"parts"`
	Owner    string `json:"owner"`
	Datetime string `json:"datetime"`
	Location string `json:"location"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "queryAirline" {
		return s.queryAirline(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordAirline" {
		return s.recordAirline(APIstub, args)
	} else if function == "queryAllAirline" {
		return s.queryAllAirline(APIstub)
	} else if function == "changeAirlineHolder" {
		return s.changeAirlineHolder(APIstub, args)
	}
	return shim.Error("Invalid SmartContract function name")
}

func (s *SmartContract) queryAirline(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("Invalid No of arguments, expecting 1")
	}
	airlineAsBytes, _ := APIstub.GetState(args[0])
	if airlineAsBytes == nil {
		return shim.Error("Could not locate AIrline")
	}
	return shim.Success(airlineAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	airline := []Airline{
		Airline{Parts: "Engine", Owner: "AirBus", Datetime: "1504054225", Location: "67.0006, -70.5476"},
		Airline{Parts: "WingTail", Owner: "IAF", Datetime: "1504057825", Location: "91.2395, -49.4594"},
		Airline{Parts: "Propelers", Owner: "Emaritaes", Datetime: "1493517025", Location: "58.0148, 59.01391"},
		Airline{Parts: "Navigation system", Owner: "Indigo", Datetime: "1496105425", Location: "-45.0945, 0.7949"},
		Airline{Parts: "Rocket Boosters", Owner: "Spacex", Datetime: "1493512301", Location: "-107.6043, 19.5003"},
		Airline{Parts: "Carrier", Owner: "Lufthansa", Datetime: "1494117101", Location: "-155.2304, -15.8723"},
		Airline{Parts: "Wheels", Owner: "PanAm", Datetime: "1496104301", Location: "103.8842, 22.1277"},
		Airline{Parts: "Landing Gear", Owner: "BritishAirways", Datetime: "1485066691", Location: "-132.3207, -34.0983"},
		Airline{Parts: "Catering Appliances", Owner: "AirAsia", Datetime: "1485153091", Location: "153.0054, 12.6429"},
		Airline{Parts: "Outfits", Owner: "VirginAirways", Datetime: "1487745091", Location: "51.9435, 8.2735"},
	}
	for i := 0; i < len(airline); i++ {
		fmt.Println("i is ", i)
		airlineAsBytes, _ := json.Marshal(airline[i])
		APIstub.PutState(strconv.Itoa(i+1), airlineAsBytes)
		fmt.Println("Added: ", airline[i])
	}
	return shim.Success(nil)
}

func (s *SmartContract) recordAirline(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 5 {
		shim.Error("Inappropriate no of Arguments, expecting 2")
	}
	var airline = Airline{Parts: args[1], Owner: args[2], Datetime: args[3], Location: args[4]}
	airlineAsBytes, _ := json.Marshal(airline)
	err := APIstub.PutState(args[0], airlineAsBytes)
	if err != nil {
		shim.Error("Failed to update Ledger")
	}
	return shim.Success(nil)
}

func (s *SmartContract) queryAllAirline(APIstub shim.ChaincodeStubInterface) sc.Response {
	startKey := "0"
	endKey := "999"
	iterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")
	arrayAlreadyWritten := false
	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			shim.Error(err.Error())
		}
		if arrayAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		arrayAlreadyWritten = true
	}
	buffer.WriteString("]")
	fmt.Println("Queried Airline Parts are : ", buffer.String())
	return shim.Success(nil)
}

func (s *SmartContract) changeAirlineHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		shim.Error("Invalid no of arguments, expected to be 1")
	}
	airlineAsBytes, _ := APIstub.GetState(args[0])
	var airline = Airline{}
	json.Unmarshal(airlineAsBytes, &airline)
	airline.Owner = args[1]
	airlineAsBytes, _ = json.Marshal(airline)
	err := APIstub.PutState(args[0], airlineAsBytes)
	if err != nil {
		shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Println("Unable to intalize the chaincode")
	}
}
