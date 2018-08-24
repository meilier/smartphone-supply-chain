package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile   string
	ChannelNames []string
	ChainCodes   []string
	initialized  bool
	UserName     string
	Secret       string
	clients      map[string]*channel.Client
	sdk          *fabsdk.FabricSDK
	event        *event.Client
	LogLevel     logging.Level
	cfg          *core.ConfigBackend
}

func (setup *FabricSetup) SetupLogLevel(lvl logging.Level) {
	if lvl != nil {
		setup.LogLevel = lvl
	}
	logging.SetLevel("fabsdk", setup.LogLevel)
	logging.SetLevel("fabsdk/common", setup.LogLevel)
	logging.SetLevel("fabsdk/fab", setup.LogLevel)
	logging.SetLevel("fabsdk/client", setup.LogLevel)
}

func (setup *FabricSetup) enrollUser() error {
	ctx := setup.sdk.Context()
	mspClient, err := msp.New(ctx)
	if err != nil {
		fmt.Printf("Failed to create msp client: %s\n", err)
		return err
	}

	_, err = mspClient.GetSigningIdentity(setup.UserName)
	if err == msp.ErrUserNotFound {
		fmt.Println("Going to enroll user")
		err = mspClient.Enroll(setup.UserName, msp.WithSecret(setup.Secret))
		if err != nil {
			fmt.Printf("Failed to enroll user: %s\n", err)
			return err
		} else {
			fmt.Printf("Success enroll user: %s\n", setup.UserName)
		}
	} else if err != nil {
		fmt.Printf("Failed to get user: %s\n", err)
		return err
	} else {
		fmt.Printf("User %s already enrolled, skip enrollment.\n", setup.UserName)
	}
	return nil
}

func (setup *FabricSetup) Initialize() error {
	if setup.initialized == true {
		return fmt.Errorf("already init")
	}
	fmt.Println("Reading connection profile..")
	c := config.FromFile(setup.ConfigFile)
	sdk, err := fabsdk.New(c)
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s\n", err)
		return err
	}
	setup.sdk = sdk
	defer setup.sdk.Close()

	setup.SetupLogLevel(setup.LogLevel)
	enrollUser(setup.sdk)

	cfg, err := setup.sdk.Config()
	if err != nil {
		fmt.Println("Failed to get sdk config. Maybe connection-profile.yaml is invalid.")
		return err
	}
	setup.cfg = cfg
	channels, ok := setup.cfg.Lookup("channels")
	if !ok {
		fmt.Println("Failed to get channels from connection-profile.yaml.")
	}

	channelsCfg, ok := channels.(map[string]interface{})

	setup.ChannelNames = make([]string, 0)
	setup.clients = make(map[string]*channel.Client)
	if ok {
		for ch := range channelsCfg {
			setup.ChannelNames = append(setup.ChannelNames, ch)
			clientChannelContext := setup.sdk.ChannelContext(ch, fabsdk.WithUser(setup.UserName))
			client, err := channel.New(clientChannelContext)
			if err != nil {
				return err
			}
			setup.clients[ch] = client
		}
	}
	setup.initialized = true
}

func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}
