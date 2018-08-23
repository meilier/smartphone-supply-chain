package main

import (
	"fmt"

	"github.com/chainHero/heroes-service/blockchain"
	"github.com/chainHero/heroes-service/web"
	"github.com/chainHero/heroes-service/web/controllers"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		user : "user1"
		secret : "arclabw401"
		channelName : "supplychannel"
		cc : "addsupplier"
		ConfigFile: "connection-profile.yaml",
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
