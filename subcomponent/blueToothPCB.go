/*add material info to ledger*/
package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MaterialChaincode struct {
}

type MaterilaInfo struct {
	Type             string `json:"Type"`
	Bandwidth		string `json:"Bandwidth"`
	Performance		 string `json:"Performance"`
	Purchaser        string `json:"Purchaser"`
}

func (t *MaterialChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	var A string
	var Aval int
	var err error

	// record num of phone
	A = "Total"
	Aval = 0

	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *MaterialChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "record" {
		return t.record(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"record\" \"query\"")
}

func (t *MaterialChaincode) record(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var B string
	var Bval int

	var record = MaterilaInfo{Type: args[1],Bandwidth:args[2], Performance:args[3], Purchaser: args[4]}

	recordbyte, _ := json.Marshal(record)
	err := stub.PutState(args[0], recordbyte)
	if err != nil {
		return shim.Error(err.Error())
	}

	B = "Total"
	Bvalbytes, err := stub.GetState(B)
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	Bval += 1

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *MaterialChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil info for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Printf("Query Response:%s\n", string(Avalbytes))
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(MaterialChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
