package tokenomics

import (
	"fmt"

	"github.com/forbole/callisto/v4/database"
	distributionsource "github.com/forbole/callisto/v4/modules/distribution/source"
	"github.com/forbole/callisto/v4/modules/dogfood"
	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/types/config"
	"github.com/imua-xyz/imuachain/utils"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the virtual tokenomics module
type Module struct {
	db          *database.Db
	cfg         *Config
	source      distributionsource.Source
	dogfoodAddr string
}

// NewModule builds a new Module instance
func NewModule(source distributionsource.Source, cfg config.Config, db *database.Db) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}
	tokenomicsCfg, err := ParseConfig(bz)
	if err != nil {
		panic(fmt.Sprintf("failed to parse tokenomic config,err:%s", err))
	}
	dogfoodCfg, err := dogfood.ParseConfig(bz)
	if err != nil {
		panic(fmt.Sprintf("failed to parse dogfood config,err:%s", err))
	}
	return &Module{
		source:      source,
		db:          db,
		cfg:         tokenomicsCfg,
		dogfoodAddr: utils.GenerateAVSAddress(dogfoodCfg.ChainIDWithoutRevision),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "tokenomics"
}
