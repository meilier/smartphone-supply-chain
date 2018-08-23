package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// QuerySupplier query the chaincode to get the state of suppliers
func (setup *FabricSetup) QuerySupplier() (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "smartisan-u2-pro-zuzhuang")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.cc, Fcn: "getSupplier", Args: [][]byte{[]byte(args[0])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}
