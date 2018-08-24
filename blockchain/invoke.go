package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//InvokeSupplier invoke addsupplier
func (setup *FabricSetup) InvokeSupplier(value []string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "addSupplier")
	args = append(args, "smartisan-u2-pro-zuzhuang")
	args = append(args, value[0])
	fmt.Println(value[0])
	args = append(args, value[1])
	fmt.Println(value[0])
	// Create a request (proposal) and send it
	response, err := setup.client.Execute(channel.Request{ChaincodeID: setup.Cc, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3])}})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	return string(response.TransactionID), nil
}
