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
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//AddLogisticsInfoChaincode chaincode implement
type AddLogisticsInfoChaincode struct {
}

type BatchInfo struct {
	Batch []string `json:"batch"`
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
	} else if function == "updateLogistics" {
		return t.updateLogistics(stub, args)
	} else if function == "deleteLogistics" {
		return t.deleteLogistics(stub, args)
	} else if function == "addBatchInfo" {
		return t.addBatchInfo(stub, args)
	} else if function == "getBatchInfo" {
		return t.getBatchInfo(stub, args)
	} else if function == "updateBatchInfo" {
		return t.updateBatchInfo(stub, args)
	} else if function == "deleteBatchInfo" {
		return t.deleteBatchInfo(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"addRecord\" \"getRecord\"")
}

// ============================================================
// // update First-Tier Suppliers information
// // key smartisan U2 pro - battery
// ============================================================

func getBatchSerial(batch string) (int, error) {
	pattern := `[\d]+`
	reg := regexp.MustCompile(pattern)
	serialString := reg.FindAllString(batch, 1)

	if len(serialString) == 0 {
		return -1, errors.New("No Serial in Batch")
	}

	serial, _ := strconv.Atoi(serialString[0])

	return serial, nil
}

func (t *AddLogisticsInfoChaincode) addBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	batch := args[1]
	batchSerial, err := getBatchSerial(batch)

	if err != nil {
		return shim.Error(err.Error())
	}

	batchInfoAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	var batchInfo BatchInfo
	json.Unmarshal(batchInfoAsBytes, &batchInfo)

	if len(batchInfo.Batch) > 0 {
		lastBatch := batchInfo.Batch[len(batchInfo.Batch)-1]
		lastSerial, _ := getBatchSerial(lastBatch)

		if lastSerial > batchSerial {
			return shim.Error("The new batchSerial should be larger than the latest batchSerial")
		}
	}

	batchInfo.Batch = append(batchInfo.Batch, batch)
	batchInfoAsBytes, _ = json.Marshal(batchInfo)
	APIstub.PutState(args[0], batchInfoAsBytes)

	return shim.Success(nil)

}

func (t *AddLogisticsInfoChaincode) getBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddLogisticsInfoChaincode) updateBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddLogisticsInfoChaincode) deleteBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddLogisticsInfoChaincode) updateLogistics(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	//args[0] key,args[1] idx,args[2] name,args[3] transit,args[4] data
	// Get the state from the ledger

	idx, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}

	transitInfoAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	var transitInfo TransitInfo
	json.Unmarshal(transitInfoAsBytes, &transitInfo)

	concretetransitinfo := transitInfo.ConcreteTransitInfo

	for i, _ := range concretetransitinfo {
		if i == idx {
			baseTransitInfo := BaseTransitInfo{Name: args[2], Transit: args[3], Manager: args[4], Date: args[5]}
			concretetransitinfo[i] = baseTransitInfo
		}
	}
	transitInfo.ConcreteTransitInfo = concretetransitinfo
	transitInfoAsBytes, _ = json.Marshal(transitInfo)
	APIstub.PutState(args[0], transitInfoAsBytes)

	return shim.Success([]byte("Successfully update"))
}

func (t *AddLogisticsInfoChaincode) deleteLogistics(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	// Get the state from the ledger
	err := APIstub.DelState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("successfully delete"))
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
