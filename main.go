package main

import (
	"fmt"
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/meilier/smartphone-supply-chain/blockchain"
	"github.com/meilier/smartphone-supply-chain/web"
	"github.com/meilier/smartphone-supply-chain/web/controllers"
)

// var (
// 	cc          = ""
// 	user        = ""
// 	secret      = ""
// 	channelName = ""
// 	lvl         = logging.INFO
// )

var (
	cc          = "sacc"
	user        = "user1"
	secret      = "arclabw401"
	channelName = "first-channel"
	lvl         = logging.INFO
)

// An implement of TargetFilter, the default filter is ledger.mspFilter which
// only pass target which in the same msp with client. But there are targets belong
// to another msp in the channel.
type AllPassTargetFilter struct {
}

func (tf *AllPassTargetFilter) Accept(peer fab.Peer) bool {
	return true
}

func queryInstalledCC(sdk *fabsdk.FabricSDK) {
	userContext := sdk.Context(fabsdk.WithUser(user))

	resClient, err := resmgmt.New(userContext)
	if err != nil {
		fmt.Println("Failed to create resmgmt: ", err)
	}

	resp2, err := resClient.QueryInstalledChaincodes()
	if err != nil {
		fmt.Println("Failed to query installed cc: ", err)
	}
	fmt.Println("Installed cc: ", resp2.GetChaincodes())
}

func queryCC(client *channel.Client, name []byte) string {
	var queryArgs = [][]byte{name}
	response, err := client.Query(channel.Request{
		ChaincodeID: cc,
		Fcn:         "query",
		Args:        queryArgs,
	})

	if err != nil {
		fmt.Println("Failed to query: ", err)
	}

	ret := string(response.Payload)
	fmt.Println("Chaincode status: ", response.ChaincodeStatus)
	fmt.Println("Payload: ", ret)
	return ret
}

func invokeCC(client *channel.Client, newValue string) {
	fmt.Println("Invoke cc with new value:", newValue)
	invokeArgs := [][]byte{[]byte("john"), []byte(newValue)}

	_, err := client.Execute(channel.Request{
		ChaincodeID: cc,
		Fcn:         "set",
		Args:        invokeArgs,
	})

	if err != nil {
		fmt.Printf("Failed to invoke: %+v\n", err)
	}
}

func readInput() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage: main.go <user-name> <user-secret> <channel> <chaincode-name>\n")
		os.Exit(1)
	}
	user = os.Args[1]
	secret = os.Args[2]
	channelName = os.Args[3]
	cc = os.Args[4]
}

func main() {

	fabricSetup := blockchain.FabricSetup{
		ConfigFile: "./connection-profile.yaml",
		UserName:   "user1",
		Secret:     "arclabw401",
		LogLevel:   logging.INFO,
	}
	fabricSetup.Initialize()
	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fabricSetup,
	}
	web.Serve(app)
}
