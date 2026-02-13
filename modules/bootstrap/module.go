package bootstrap

import (
	"context"
	"fmt"

	"github.com/xrpscan/xrpl-go"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	callistodb "github.com/forbole/callisto/v4/database"
	"github.com/forbole/callisto/v4/modules/bootstrap/bootstrap_binding"
	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/types/config"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	database            *callistodb.Db
	EthHTTPClient       *ethclient.Client
	EthWSClient         *ethclient.Client
	FeederEthHTTPClient *ethclient.Client
	XrpClient           *xrpl.Client
	Config              Config
	BootstrapAddr       common.Address
	// No dedicated clients for BTC and XRP, since we may call their RPCs
	// directly using the http package.
	ctx               context.Context
	bootstrapSession  *bootstrap_binding.BootstrapCallerSession
	bootstrapFilterer *bootstrap_binding.BootstrapFilterer
	// Address mappings for 1-1 binding validation
	btcAddressMappings map[string]string // bitcoin -> imuachain
	xrpAddressMappings map[string]string // xrp -> imuachain
}

// NewModule builds a new Module instance
func NewModule(
	cfg config.Config,
	database *callistodb.Db,
) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}

	bootstrapCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}
	if !common.IsHexAddress(bootstrapCfg.BootstrapAddr) {
		panic(fmt.Sprintf("invalid bootstrap address:%s", bootstrapCfg.BootstrapAddr))
	}
	bootstrapAddr := common.HexToAddress(bootstrapCfg.BootstrapAddr)

	httpRC, err := rpc.DialContext(context.Background(), bootstrapCfg.ETHHttp)
	if err != nil {
		panic(err)
	}
	ethHTTPClient := ethclient.NewClient(httpRC)

	websocketRC, err := rpc.DialContext(context.Background(), bootstrapCfg.ETHWebsocket)
	if err != nil {
		panic(err)
	}
	ethWSClient := ethclient.NewClient(websocketRC)

	feederRC, err := rpc.DialContext(context.Background(), bootstrapCfg.FeederEthHTTP)
	if err != nil {
		panic(err)
	}

	xrpClient := xrpl.NewClient(xrpl.ClientConfig{URL: bootstrapCfg.XRPRPC})
	err = xrpClient.Ping([]byte("PING"))
	if err != nil {
		panic(fmt.Errorf("failed to ping XRP client at %s: %w", bootstrapCfg.XRPRPC, err))
	}

	// create the sessions for bootstrap and storage contracts.
	ctx := context.Background()
	bootstrapCaller, err := bootstrap_binding.NewBootstrapCaller(bootstrapAddr, ethHTTPClient)
	if err != nil {
		panic(fmt.Errorf("failed to new bootstrap caller,err:%s", err))
	}
	bootstrapSession := &bootstrap_binding.BootstrapCallerSession{
		Contract: bootstrapCaller,
		CallOpts: bind.CallOpts{Context: ctx},
	}

	// create the filterer to subscribe all related events
	bootstrapFilterer, err := bootstrap_binding.NewBootstrapFilterer(bootstrapAddr, ethWSClient)
	if err != nil {
		panic(fmt.Errorf("failed to new bootstrap filterer,err:%s", err))
	}

	module := &Module{
		database:            database,
		EthHTTPClient:       ethHTTPClient,
		EthWSClient:         ethWSClient,
		FeederEthHTTPClient: ethclient.NewClient(feederRC),
		XrpClient:           xrpClient,
		Config:              *bootstrapCfg,
		BootstrapAddr:       bootstrapAddr,
		ctx:                 ctx,
		bootstrapSession:    bootstrapSession,
		bootstrapFilterer:   bootstrapFilterer,
		btcAddressMappings:  make(map[string]string),
		xrpAddressMappings:  make(map[string]string),
	}
	return module
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bootstrap"
}
