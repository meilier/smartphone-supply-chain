package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// QueryCC query the chaincode to get the state of suppliers
func (setup *FabricSetup) QueryCC(channelName, chainCode, fcn string, key []byte) (string, error) {
	var queryArgs = [][]byte{key}
	fmt.Println("why a is", setup.clients)
	client := setup.clients[channelName]
	fmt.Println("key is ", string(key))
	fmt.Println("queryArgs is", queryArgs)
	fmt.Println("client is", client)
	response, err := client.Query(channel.Request{ChaincodeID: chainCode, Fcn: fcn, Args: queryArgs})
	fmt.Println(err)
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}
	return string(response.Payload), nil

}
