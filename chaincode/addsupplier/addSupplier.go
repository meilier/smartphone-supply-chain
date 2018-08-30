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
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//AddSupplierChaincode chaincode implement
type AddSupplierChaincode struct {
}

type BatchInfo struct {
	Batch []string `json:"batch"`
}

//CompanyInfo define the company structure, with x properties.  Structure tags are used by encoding/json library
type CompanyInfo struct {
	Name          string             `json:"name"`
	Location      string             `json:"location"`
	ComponentInfo string             `json:"componentinfo"`
	Subcomponent  []SubComponentInfo `json:"subcomponent"`
}

//
type SubComponentInfo struct {
	SubName        string
	SubCompanyName string
	SubLocation    string
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
	} else if function == "updateSupplier" {
		return t.updateSupplier(stub, args)
	} else if function == "deleteSupplier" {
		return t.deleteSupplier(stub, args)
	} else if function == "addBatchInfo" {
		return t.addBatchInfo(stub, args)
	} else if function == "getBatchInfo" {
		return t.getBatchInfo(stub, args)
	} else if function == "updateBatchInfo" {
		return t.updateBatchInfo(stub, args)
	} else if function == "deleteBatchInfo" {
		return t.deleteBatchInfo(stub, args)
	} else if function == "addCompanyInfo" {
		return t.addCompanyInfo(stub, args)
	} else if function == "getCompanyInfo" {
		return t.getCompanyInfo(stub, args)
	} else if function == "updateCompanyInfo" {
		return t.updateCompanyInfo(stub, args)
	} else if function == "deleteCompanyInfo" {
		return t.deleteCompanyInfo(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"addRecord\" \"getRecord\"")
}

// ============================================================
// // update First-Tier Suppliers information
// // key smartisan U2 pro - battery
// ============================================================
func (t *AddSupplierChaincode) addBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddSupplierChaincode) getBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddSupplierChaincode) updateBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddSupplierChaincode) deleteBatchInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddSupplierChaincode) addSupplier(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	var ccompany CompanyInfo
	var record = SubComponentInfo{SubName: args[1], SubCompanyName: args[2], SubLocation: args[3]}

	// Get the state from the ledger first
	ccompanyBeforAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	json.Unmarshal(ccompanyBeforAsBytes, &ccompany)

	ccompany.Subcomponent = append(ccompany.Subcomponent, record)

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
	var ccompany CompanyInfo
	json.Unmarshal(ccompanyAsBytes, &ccompany)

	if err != nil {
		return shim.Error(err.Error())
	}

	supplierInfo := ccompany.Subcomponent
	supplierInfoAsbytes, _ := json.Marshal(supplierInfo)

	return shim.Success(supplierInfoAsbytes)
}

func (t *AddSupplierChaincode) updateSupplier(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	ccompanyAsBytes, err := APIstub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	var ccompany CompanyInfo
	json.Unmarshal(ccompanyAsBytes, &ccompany)
	supplier := ccompany.Subcomponent

	var newSupplier []SubComponentInfo
	for _, v := range supplier {
		if v.SubName != args[1] {
			newSupplier = append(newSupplier, v)
		} else {
			newSub := SubComponentInfo{SubName: args[1], SubCompanyName: args[2], SubLocation: args[3]}
			newSupplier = append(newSupplier, newSub)
		}
	}

	ccompany.Subcomponent = newSupplier
	ccompanyAsBytes, _ = json.Marshal(ccompany)
	APIstub.PutState(args[0], ccompanyAsBytes)

	return shim.Success([]byte("success delete"))

	return shim.Success(ccompanyAsBytes)
}

func (t *AddSupplierChaincode) deleteSupplier(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	//args[0] key,args[1] subName
	// Get the state from the ledger
	ccompanyAsBytes, err := APIstub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	var ccompany CompanyInfo
	json.Unmarshal(ccompanyAsBytes, &ccompany)
	supplier := ccompany.Subcomponent

	var newSupplier []SubComponentInfo
	for _, v := range supplier {
		if v.SubName != args[1] {
			newSupplier = append(newSupplier, v)
		}
	}

	ccompany.Subcomponent = newSupplier
	ccompanyAsBytes, _ = json.Marshal(ccompany)
	APIstub.PutState(args[0], ccompanyAsBytes)

	return shim.Success([]byte("successfuly delete"))
}

func (t *AddSupplierChaincode) addCompanyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	cc := CompanyInfo{Name: args[1], Location: args[2], ComponentInfo: args[3]}
	var ccompany CompanyInfo
	ccompanyBeforAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	json.Unmarshal(ccompanyBeforAsBytes, &ccompany)
	cc.Subcomponent = ccompany.Subcomponent

	ccompanyAsBytes, _ := json.Marshal(cc)
	APIstub.PutState(args[0], ccompanyAsBytes)

	return shim.Success(nil)

}

func (t *AddSupplierChaincode) getCompanyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

func (t *AddSupplierChaincode) updateCompanyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	cc := CompanyInfo{Name: args[1], Location: args[2], ComponentInfo: args[3]}
	var ccompany CompanyInfo
	ccompanyBeforAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	json.Unmarshal(ccompanyBeforAsBytes, &ccompany)
	cc.Subcomponent = ccompany.Subcomponent

	ccompanyAsBytes, _ := json.Marshal(cc)
	APIstub.PutState(args[0], ccompanyAsBytes)

	return shim.Success(nil)
}

func (t *AddSupplierChaincode) deleteCompanyInfo(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
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

// ======================================AddSupplierChaincode=============================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(AddSupplierChaincode))
	if err != nil {
		fmt.Printf("Error starting AddSupplier chaincode: %s", err)
	}
}
