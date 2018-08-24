package main

import (
	"fmt"

	"github.com/meilier/smartphone-supply-chain/blockchain"
	"github.com/meilier/smartphone-supply-chain/web"
	"github.com/meilier/smartphone-supply-chain/web/controllers"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		User:        "user1",
		Secret:      "arclabw401",
		ChannelName: "supplychannel",
		Cc:          "addsupplier",
		ConfigFile:  "connection-profile.yaml",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	defer fSetup.CloseSDK()

	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}
