package types

import (
	"fmt"
	"os"

	"github.com/imua-xyz/imuachain/x/feedistribution/types"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/forbole/juno/v5/node/remote"
	"github.com/forbole/juno/v5/types/params"

	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	delegationtypes "github.com/imua-xyz/imuachain/x/delegation/types"
	dogfoodkeeper "github.com/imua-xyz/imuachain/x/dogfood/keeper"
	dogfoodtypes "github.com/imua-xyz/imuachain/x/dogfood/types"
	epochstypes "github.com/imua-xyz/imuachain/x/epochs/types"
	imminttypes "github.com/imua-xyz/imuachain/x/immint/types"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/forbole/juno/v5/node/local"

	nodeconfig "github.com/forbole/juno/v5/node/config"

	banksource "github.com/forbole/callisto/v4/modules/bank/source"
	localbanksource "github.com/forbole/callisto/v4/modules/bank/source/local"
	remotebanksource "github.com/forbole/callisto/v4/modules/bank/source/remote"

	epochssource "github.com/forbole/callisto/v4/modules/epochs/source"
	localepochssource "github.com/forbole/callisto/v4/modules/epochs/source/local"
	remoteepochssource "github.com/forbole/callisto/v4/modules/epochs/source/remote"

	slashingsource "github.com/forbole/callisto/v4/modules/slashing/source"
	localslashingsource "github.com/forbole/callisto/v4/modules/slashing/source/local"
	remoteslashingsource "github.com/forbole/callisto/v4/modules/slashing/source/remote"

	immintsource "github.com/forbole/callisto/v4/modules/immint/source"
	localimmintsource "github.com/forbole/callisto/v4/modules/immint/source/local"
	remoteimmintsource "github.com/forbole/callisto/v4/modules/immint/source/remote"

	assetssource "github.com/forbole/callisto/v4/modules/assets/source"
	localassetssource "github.com/forbole/callisto/v4/modules/assets/source/local"
	remoteassetssource "github.com/forbole/callisto/v4/modules/assets/source/remote"

	delegationsource "github.com/forbole/callisto/v4/modules/delegation/source"
	localdelegationsource "github.com/forbole/callisto/v4/modules/delegation/source/local"
	remotedelegationsource "github.com/forbole/callisto/v4/modules/delegation/source/remote"

	dogfoodsource "github.com/forbole/callisto/v4/modules/dogfood/source"
	localdogfoodsource "github.com/forbole/callisto/v4/modules/dogfood/source/local"
	remotedogfoodsource "github.com/forbole/callisto/v4/modules/dogfood/source/remote"

	oraclesource "github.com/forbole/callisto/v4/modules/oracle/source"
	localoraclesource "github.com/forbole/callisto/v4/modules/oracle/source/local"
	remoteoraclesource "github.com/forbole/callisto/v4/modules/oracle/source/remote"

	distributionsource "github.com/forbole/callisto/v4/modules/distribution/source"
	localdistributionsource "github.com/forbole/callisto/v4/modules/distribution/source/local"
	remotedistributionsource "github.com/forbole/callisto/v4/modules/distribution/source/remote"

	imuachainapp "github.com/imua-xyz/imuachain/app"
)

type Sources struct {
	BankSource banksource.Source
	// DistrSource    distrsource.Source
	// GovSource      govsource.Source
	// MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	// StakingSource  stakingsource.Source
	EpochsSource       epochssource.Source
	ImmintSource       immintsource.Source
	AssetsSource       assetssource.Source
	DelegationSource   delegationsource.Source
	DogfoodSource      dogfoodsource.Source
	OracleSource       oraclesource.Source
	DistributionSource distributionsource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	app := imuachainapp.NewImuachainApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, nil, cfg.Home,
		0, simappparams.EncodingConfig{}, nil, nil,
	)

	sources := &Sources{
		BankSource: localbanksource.NewSource(source, banktypes.QueryServer(app.BankKeeper)),
		// DistrSource:    localdistrsource.NewSource(source, distrtypes.QueryServer(app.DistrKeeper)),
		// GovSource:      localgovsource.NewSource(source, govtypesv1.QueryServer(app.GovKeeper)),
		// MintSource:     localmintsource.NewSource(source, minttypes.QueryServer(app.MintKeeper)),
		SlashingSource: localslashingsource.NewSource(source, slashingtypes.QueryServer(app.SlashingKeeper)),
		EpochsSource:   localepochssource.NewSource(source, epochstypes.QueryServer(app.EpochsKeeper)),
		ImmintSource:   localimmintsource.NewSource(source, imminttypes.QueryServer(app.ImmintKeeper)),
		// StakingSource:  localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: app.StakingKeeper}),
		AssetsSource:       localassetssource.NewSource(source, assetstypes.QueryServer(app.AssetsKeeper)),
		DelegationSource:   localdelegationsource.NewSource(source, delegationtypes.QueryServer(&app.DelegationKeeper)),
		DogfoodSource:      localdogfoodsource.NewSource(source, dogfoodkeeper.NewQueryServer(app.StakingKeeper)),
		OracleSource:       localoraclesource.NewSource(source, oracletypes.QueryServer(app.OracleKeeper)),
		DistributionSource: localdistributionsource.NewSource(source, types.QueryServer(app.DistrKeeper)),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(app, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(app, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(app, "memKeys")
	if err != nil {
		return nil, err
	}

	err = source.InitStores()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource: remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		// DistrSource:    remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		// GovSource:      remotegovsource.NewSource(source, govtypesv1.NewQueryClient(source.GrpcConn)),
		// MintSource:     remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource: remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		// StakingSource:  remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),
		EpochsSource:       remoteepochssource.NewSource(source, epochstypes.NewQueryClient(source.GrpcConn)),
		ImmintSource:       remoteimmintsource.NewSource(source, imminttypes.NewQueryClient(source.GrpcConn)),
		AssetsSource:       remoteassetssource.NewSource(source, assetstypes.NewQueryClient(source.GrpcConn)),
		DelegationSource:   remotedelegationsource.NewSource(source, delegationtypes.NewQueryClient(source.GrpcConn)),
		DogfoodSource:      remotedogfoodsource.NewSource(source, dogfoodtypes.NewQueryClient(source.GrpcConn)),
		OracleSource:       remoteoraclesource.NewSource(source, oracletypes.NewQueryClient(source.GrpcConn)),
		DistributionSource: remotedistributionsource.NewSource(source, types.NewQueryClient(source.GrpcConn)),
	}, nil
}
