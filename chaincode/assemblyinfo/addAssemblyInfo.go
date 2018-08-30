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
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//AddAssemblyInfoChaincode chaincode implement
type AddAssemblyInfoChaincode struct {
}

type BatchInfo struct {
	Batch []string `json:"batch"`
}

//AssemblyInfo define basic assembly information in factory
type AssemblyInfo struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Manager  string `json:"manager"`
	Date     string `json:"date"`
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
	} else if function == "deleteAssemblyInfo" {
		return t.deleteAssemblyInfo(stub, args)
	} else if function == "updateAssemblyInfo" {
		return t.updateAssemblyInfo(stub, args)
	} else if function == "addBatchInfo" {
		return t.addBatchInfo(stub, args)
	} else if function == "getBatchInfo" {
		return t.getBatchInfo(stub, args)
	} else if function == "deleteBatchInfo" {
		return t.deleteBatchInfo(stub, args)
	} else if function == "updateBatchInfo" {
		return t.updateBatchInfo(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"addAssemblyInfo\" \"getAssemblyInfo\"")
}

// ============================================================
// // update First-Tier Suppliers information
// // key smartisan U2 pro - battery
// ============================================================
func (t *AddAssemblyInfoChaincode) addAssemblyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var assemblyinfo = AssemblyInfo{Name: args[1], Location: args[2], Manager: args[3], Date: args[4]}

	assemblyInfoAsBytes, _ := json.Marshal(assemblyinfo)
	APIstub.PutState(args[0], assemblyInfoAsBytes)

	return shim.Success(nil)

}

func (t *AddAssemblyInfoChaincode) deleteAssemblyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	err := APIstub.DelState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *AddAssemblyInfoChaincode) updateAssemblyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var assemblyinfo = AssemblyInfo{Name: args[1], Location: args[2], Manager: args[3], Date: args[4]}

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
func (t *AddAssemblyInfoChaincode) addBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	var batchInfo BatchInfo
	batch := args[1]
	batchInfoAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	json.Unmarshal(batchInfoAsBytes, &batchInfo)
	batchInfo.Batch = append(batchInfo.Batch, batch)
	batchInfoAsBytes, _ = json.Marshal(batchInfo)
	APIstub.PutState(args[0], batchInfoAsBytes)

	return shim.Success(nil)

}

func (t *AddAssemblyInfoChaincode) getBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Get the state from the ledger
	batchInfoAsBytes, err := APIstub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(batchInfoAsBytes)
}

func (t *AddAssemblyInfoChaincode) updateBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	//args[0] key , args[1] idx , args[2] operation ,args[4] value

	batchInfoAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	idx, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	var batchInfo BatchInfo
	json.Unmarshal(batchInfoAsBytes, &batchInfo)

	batch := batchInfo.Batch

	fmt.Println(batch)

	if args[2] == "change" {
		for i, _ := range batch {
			if i == idx {
				batch[i] = args[3]
			}
		}
	} else if args[2] == "delete" {
		var newBatch []string
		for i, v := range batch {
			if i == idx {
				continue
			} else {
				newBatch = append(newBatch, v)
			}
			batch = newBatch
		}
	} else if args[2] == "insert" {
		var newBatch []string
		for i, v := range batch {
			if i == idx {
				newBatch = append(newBatch, args[3])
			}
			newBatch = append(newBatch, v)
		}
		batch = newBatch
	} else {
		return shim.Error("Invaild Args[2] , Expecting ('change','delete','insert') ")
	}

	batchInfo = BatchInfo{Batch: batch}
	batchInfoAsBytes, _ = json.Marshal(batchInfo)

	APIstub.PutState(args[0], batchInfoAsBytes)

	return shim.Success([]byte("Update successfully"))
}

func (t *AddAssemblyInfoChaincode) deleteBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Get the state from the ledger
	err := APIstub.DelState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Delete successfully"))
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
