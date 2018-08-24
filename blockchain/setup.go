package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/pkg/errors"
)

// FabricSetup implementation
type FabricSetup struct {
	User        string
	Secret      string
	ChannelName string
	Cc          string
	ConfigFile  string
	initialized bool
	client      *channel.Client
	sdk         *fabsdk.FabricSDK
}

// Initialize reads the configuration file and sets up the client, chain and event hub
func (setup *FabricSetup) Initialize() error {

	// Add parameters for the initialization
	if setup.initialized {
		return errors.New("sdk already initialized")
	}

	// Initialize the SDK with the configuration file
	fmt.Println("Reading connection profile..")
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "failed to create SDK")
	}
	setup.sdk = sdk

	setup.setupLogLevel()
	setup.enrollUser(sdk)

	clientChannelContext := setup.sdk.ChannelContext(setup.ChannelName, fabsdk.WithUser(setup.User))

	fmt.Println("\n====== Chaincode =========")

	setup.client, _ = channel.New(clientChannelContext)

	fmt.Println("SDK created")
	fmt.Println("Initialization Successful")
	setup.initialized = true
	return nil
}

func (setup *FabricSetup) enrollUser(sdk *fabsdk.FabricSDK) {
	ctx := sdk.Context()
	mspClient, err := msp.New(ctx)
	if err != nil {
		fmt.Printf("Failed to create msp client: %s\n", err)
	}

	_, err = mspClient.GetSigningIdentity(setup.User)
	if err == msp.ErrUserNotFound {
		fmt.Println("Going to enroll user")
		err = mspClient.Enroll(setup.User, msp.WithSecret(setup.Secret))

		if err != nil {
			fmt.Printf("Failed to enroll user: %s\n", err)
		} else {
			fmt.Printf("Success enroll user: %s\n", setup.User)
		}

	} else if err != nil {
		fmt.Printf("Failed to get user: %s\n", err)
	} else {
		fmt.Printf("User %s already enrolled, skip enrollment.\n", setup.User)
	}
}

func (setup *FabricSetup) setupLogLevel() {
	logging.SetLevel("fabsdk", logging.INFO)
	logging.SetLevel("fabsdk/common", logging.INFO)
	logging.SetLevel("fabsdk/fab", logging.INFO)
	logging.SetLevel("fabsdk/client", logging.INFO)
}

//CloseSDK if app is closing
func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}
