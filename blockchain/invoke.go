package blockchain

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

//InvokeCC add info to chain
func (setup *FabricSetup) InvokeCC(channelName, chainCode, fcn string, newValue []string) (string, error) {
	fmt.Println("Invoke cc with new value:", newValue)
	client := setup.clients[channelName]
	// Prepare arguments
	invokeArgs := [][]byte{[]byte(strings.Join(newValue, ""))}
	// Create a request (proposal) and send it
	fmt.Println("parameter is ", channelName, chainCode, fcn, newValue)
	fmt.Println("len is", len(invokeArgs))
	response, err := client.Execute(channel.Request{ChaincodeID: chainCode, Fcn: fcn, Args: invokeArgs})
	fmt.Println("err is", err)
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}
	return string(response.TransactionID), nil

}
