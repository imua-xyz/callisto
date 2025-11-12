package distribution

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// updateCommunityPool fetch total amount of coins in the system from RPC and store it into database
func (m *Module) updateCommunityPool(height int64) error {
	log.Debug().Str("module", "distribution").Int64("height", height).Msg("getting and updating the community pools for all AVSs")

	allAVSs, err := m.db.GetAllAvsAddrs()
	if err != nil {
		return fmt.Errorf("error while getting all AVSs: %s", err)
	}
	for _, avs := range allAVSs {
		pool, err := m.source.AVSCommunityPool(height, avs)
		if err != nil {
			return fmt.Errorf("error while getting avs community pool, avs:%s, err:%s", avs, err)
		}
		err = m.db.SaveCommunityPool(avs, pool, height)
		if err != nil {
			return fmt.Errorf("error while saving community pool for avs: %s,err:%s", avs, err)
		}
	}
	return nil
}
