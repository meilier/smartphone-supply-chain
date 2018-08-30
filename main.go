package main

import (
	"fmt"

	"github.com/meilier/smartphone-supply-chain/web/webutil"

	"github.com/meilier/smartphone-supply-chain/blockchain"
	"github.com/meilier/smartphone-supply-chain/web"
	"github.com/meilier/smartphone-supply-chain/web/controllers"
)

var fabricSetups map[string]*blockchain.FabricSetup

func main() {
	fabricSetups = make(map[string]*blockchain.FabricSetup)
	// Definition of the Fabric SDK properties
	//for k, am := range webutil.Orgnization {
	for _, am := range webutil.Orgnization {
		//if k == "sales" || k == "smartphone" {
		for _, user := range am {
			fabricSetups[user.UserName] = &blockchain.FabricSetup{
				User:       user.UserName,
				Secret:     user.Secret,
				ConfigFile: user.FilePath,
			}
			// Initialization of the Fabric SDK from the previously set properties
			err := fabricSetups[user.UserName].Initialize()
			if err != nil {
				fmt.Printf("Unable to initialize the sFabric SDK: %v\n", err)
				return
			}
		}

		//}
	}

	// Close SDK
	for _, v := range fabricSetups {
		defer v.CloseSDK()
	}

	// Launch the web application listening
	app := &controllers.Application{
		Fabric: fabricSetups,
	}
	web.Serve(app)
}
