/*
This chaincode is implemented to add assembly info to ledger.
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

//AddAssemblyInfoChaincode chaincode implement
type AddAssemblyInfoChaincode struct {
}

//AssemblyInfo define basic assembly information in factory
type AssemblyInfo struct {
	Name     string `json:"name"`
	Deliver string `json:"deliver"`
	Date  string `json:"date"`
}

//Init init fuction
func (t *AddAssemblyInfoChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke function
func (t *AddAssemblyInfoChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	// Retrieve the requested Smart Contract function , arguments and transient
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "addAssemblyInfo" {
		return t.addAssemblyInfo(stub, args)
	} else if function == "getAssemblyInfo" {
		return t.getAssemblyInfo(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"addRecord\" \"getRecord\"")
}

// ============================================================
// // update First-Tier Suppliers information
// // key smartisan U2 pro - battery
// ============================================================
func (t *AddAssemblyInfoChaincode) addAssemblyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	var assemblyinfo = AssemblyInfo{Name: args[1], Deliver: args[2], Date: args[3]}

	assemblyInfoAsBytes, _ := json.Marshal(assemblyinfo)
	APIstub.PutState(args[0], assemblyInfoAsBytes)

	return shim.Success(nil)

}

func (t *AddAssemblyInfoChaincode) getAssemblyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Get the state from the ledger
	assemblyInfoAsBytes, err := APIstub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(assemblyInfoAsBytes)
}

// ======================================AddAssemblyInfoChaincode=============================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(AddAssemblyInfoChaincode))
	if err != nil {
		fmt.Printf("Error starting Consensual Letter chaincode: %s", err)
	}
}
