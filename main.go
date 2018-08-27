package main

import (
	"fmt"

	"github.com/meilier/smartphone-supply-chain/blockchain"
	"github.com/meilier/smartphone-supply-chain/web"
	"github.com/meilier/smartphone-supply-chain/web/controllers"
)

var Accout map[string]string

var AphoneSerialNumber = [5]string{"10000000", "10000001", "10000002", "10000003", "10000004"}
var XphoneSerialNumber = [5]string{"20000000", "20000001", "20000002", "20000003", "20000004"}

type CodeInfo struct {
	User        string
	ChannelName string
	CodeName    string
	FilePath    string
}

var Orgnization map[string][]CodeInfo

func init() {
	Accout = make(map[string]string)
	Orgnization = make(map[string][]CodeInfo)
	Accout["wzx"] = "arclabw401wzx"
	Accout["lwh"] = "arclabw401lwh"
	Accout["wyh"] = "arclabw401wyh"
	Accout["yzx"] = "arclabw401yzx"
	Accout["xjx"] = "arclabw401xjx"
	// smsc := CodeInfo{"wzx", "supplychannel", "addsupplier", "./profile/smartphone/connection-profile-wzx.yaml"}
	// smac := CodeInfo{"wzx", "assemblychannel", "assembly", "./profile/smartphone/connection-profile-wzx.yaml"}
	// Orgnization["smartphone"] = append(Orgnization["smartphone"], smsc)
	// Orgnization["smartphone"] = append(Orgnization["smartphone"], smac)

	// su := CodeInfo{"lwh", "supplychannel", "addsupplier", "./profile/supplier/connection-profile-lwh.yaml"}
	// Orgnization["smartphone"] = append(Orgnization["smartphone"], smsc)
	// as := CodeInfo{"wyh", "supplychannel", "addsupplier", "./profile/assembly/connection-profile-wyh.yaml"}
	// Orgnization["smartphone"] = append(Orgnization["smartphone"], smsc)
	// lo := CodeInfo{"yzx", "supplychannel", "addsupplier", "./profile/logistics/connection-profile-yzx.yaml"}
	// Orgnization["smartphone"] = append(Orgnization["smartphone"], smsc)
	// st := CodeInfo{"xjx", "supplychannel", "addsupplier", "./profile/store/connection-profile-xjx.yaml"}
	// Orgnization["smartphone"] = append(Orgnization["smartphone"], smsc)

}

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		User:        "wzx",
		Secret:      "arclabw401wzx",
		ChannelName: "supplychannel",
		Cc:          "addsupplier",
		ConfigFile:  "./profile/smartphone/connection-profile-wzx.yaml",
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
