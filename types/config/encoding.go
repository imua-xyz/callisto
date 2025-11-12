package config

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/evmos/evmos/v16/encoding/codec"
	"github.com/forbole/juno/v5/types/params"
)

// MakeEncodingConfig creates an EncodingConfig to properly handle all the messages
func MakeEncodingConfig(managers []module.BasicManager) func() params.EncodingConfig {
	return func() params.EncodingConfig {
		encodingConfig := params.MakeTestEncodingConfig()
		manager := mergeBasicManagers(managers)
		manager.RegisterLegacyAminoCodec(encodingConfig.Amino)
		manager.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		// register evmos codec - which contains the sdk codec
		codec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		codec.RegisterLegacyAminoCodec(encodingConfig.Amino)
		return encodingConfig
	}
}

// mergeBasicManagers merges the given managers into a single module.BasicManager
func mergeBasicManagers(managers []module.BasicManager) module.BasicManager {
	union := module.BasicManager{}
	for _, manager := range managers {
		for k, v := range manager {
			union[k] = v
		}
	}
	return union
}
