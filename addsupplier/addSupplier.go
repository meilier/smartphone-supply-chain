/*
This chaincode is implemented to add First-tier suppliers to ledger.
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

//AddSupplierChaincode chaincode implement
type AddSupplierChaincode struct {
}

//CompanyInfo define the company structure, with x properties.  Structure tags are used by encoding/json library
type CompanyInfo struct {
	ConcreteCompanyInfo []BaseCompanyInfo `json:"concretecompanyinfo"`
}

//BaseCompanyInfo define basic company info
type BaseCompanyInfo struct {
	Name     string `json:"name"`
	Location string `json:"year"`
}

//Init init fuction
func (t *AddSupplierChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke function
func (t *AddSupplierChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	// Retrieve the requested Smart Contract function , arguments and transient
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "addSupplier" {
		return t.addSupplier(stub, args)
	} else if function == "getSupplier" {
		return t.getSupplier(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"addRecord\" \"getRecord\"")
}

// ============================================================
// // update First-Tier Suppliers information
// // key smartisan U2 pro - battery
// ============================================================
func (t *AddSupplierChaincode) addSupplier(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	var ccompany CompanyInfo
	var record = BaseCompanyInfo{Name: args[1], Location: args[2]}
	ccompany.ConcreteCompanyInfo = append(ccompany.ConcreteCompanyInfo, record)

	ccompanyAsBytes, _ := json.Marshal(ccompany)
	APIstub.PutState(args[0], ccompanyAsBytes)

	return shim.Success(nil)

}

func (t *AddSupplierChaincode) getSupplier(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Get the state from the ledger
	ccompanyAsBytes, err := APIstub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(ccompanyAsBytes)
}

// ======================================PersonalRecordChaincode=============================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(AddSupplierChaincode))
	if err != nil {
		fmt.Printf("Error starting Consensual Letter chaincode: %s", err)
	}
}
