package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// QueryCC query the chaincode to get the state of suppliers
func (setup *FabricSetup) DeleteCC(channelName, chainCode, fcn string, newValue []string) (string, error) {
	invokeArgs := [][]byte{}
	for _, s := range newValue {
		invokeArgs = append(invokeArgs, []byte(s))
	}
	fmt.Println("why a is", setup.clients)
	client := setup.clients[channelName]
	fmt.Println("newvalue is ", newValue)
	fmt.Println("client is", client)
	response, err := client.Execute(channel.Request{ChaincodeID: chainCode, Fcn: fcn, Args: invokeArgs})
	fmt.Println(err)
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil

}
