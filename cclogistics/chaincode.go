/*
This chaincode is implemented to add personal record to ledger. Add record or encrypted
record to the ledger and get record or encrypted record form the ledger.
*/

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// "github.com/hyperledger/fabric/core/chaincode/shim/ext/entities"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// const DECKEY = "DECKEY"
// const ENCKEY = "ENCKEY"
// const IV = "IV"

// Personal Record Chaincode Implementation
type PersonalRecordChaincode struct {
	bccspInst bccsp.BCCSP
}

// Define the record structure, with 4 properties.  Structure tags are used by encoding/json library
type Record struct {
	Date     string `json:"date"`
	Position string `json:"position"`
	Name		 string `json:"name"`
}

// init fuction
func (t *PersonalRecordChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *PersonalRecordChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	// Retrieve the requested Smart Contract function , arguments and transient
	function, args := stub.GetFunctionAndParameters()
	// tMap, err := stub.GetTransient()
	// if err != nil {
	// 	return shim.Error(fmt.Sprintf("Could not retrieve transient, err %s", err))
	// }
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "addRecord" {
		return t.addRecord(stub, args)
	} else if function == "getRecord" {
		return t.getRecord(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"addRecord\" \"getRecord\"")
}

// ============================================================
// addRecoed - create a new record, store into chaincode state
// ============================================================
func (t *PersonalRecordChaincode) addRecord(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	var record = Record{Date: args[1], Position: args[2], Name: args[3]}

	recordAsBytes, _ := json.Marshal(record)
	APIstub.PutState(args[0], recordAsBytes)

	return shim.Success(nil)

}

func (t *PersonalRecordChaincode) getRecord(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	resultsIterator, err := APIstub.GetHistoryForKey(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if response.IsDelete {
			continue
		}
		recordAsBytes := response.Value
		record := Record{}

		json.Unmarshal(recordAsBytes, &record)
		if record.Date == args[1] {
			return shim.Success([]byte(record.Position))
		} else {
			continue
		}

	}

	return shim.Success(nil)
}

// ======================================PersonalRecordChaincode=============================================
// Main
// ===================================================================================
func main() {
	factory.InitFactories(nil)

	err := shim.Start(&PersonalRecordChaincode{factory.GetDefault()})
	if err != nil {
		fmt.Printf("Error starting PersonalRecordChaincode chaincode: %s", err)
	}
}
