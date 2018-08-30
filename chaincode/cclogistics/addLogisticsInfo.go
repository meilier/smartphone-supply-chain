/*
This chaincode is implemented to add Logistics to ledger.
*/

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//AddLogisticsInfoChaincode chaincode implement
type AddLogisticsInfoChaincode struct {
}

//TransitInfo define the transit structure, with x properties.  Structure tags are used by encoding/json library
type TransitInfo struct {
	ConcreteTransitInfo []BaseTransitInfo `json:"concretetransitinfo"`
}

//BaseTransitInfo define basic transit info
type BaseTransitInfo struct {
	Name    string `json:"name"`
	Transit string `json:"transit"`
	Manager string `json:"manager"`
	Date    string `json:"date"`
}

//Init init fuction
func (t *AddLogisticsInfoChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke function
func (t *AddLogisticsInfoChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	// Retrieve the requested Smart Contract function , arguments and transient
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "addLogistics" {
		return t.addLogistics(stub, args)
	} else if function == "getLogistics" {
		return t.getLogistics(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"addRecord\" \"getRecord\"")
}

// ============================================================
// // update First-Tier Suppliers information
// // key smartisan U2 pro - battery
// ============================================================
func (t *AddLogisticsInfoChaincode) addLogistics(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var transitInfoBefor TransitInfo
	var bTransit = BaseTransitInfo{Name: args[1], Transit: args[2], Manager: args[3], Date: args[4]}

	// Get the state from the ledger first
	transitInfoBeforAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	json.Unmarshal(transitInfoBeforAsBytes, &transitInfoBefor)

	transitInfoBefor.ConcreteTransitInfo = append(transitInfoBefor.ConcreteTransitInfo, bTransit)

	transitInfoAsBytes, _ := json.Marshal(transitInfoBefor)
	APIstub.PutState(args[0], transitInfoAsBytes)

	return shim.Success(nil)

}

func (t *AddLogisticsInfoChaincode) getLogistics(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Get the state from the ledger
	transitInfoAsBytes, err := APIstub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(transitInfoAsBytes)
}

// ======================================AddLogisticsInfoChaincode=============================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(AddLogisticsInfoChaincode))
	if err != nil {
		fmt.Printf("Error starting AddLogisticsInfo chaincode: %s", err)
	}
}
