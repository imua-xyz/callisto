package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"
	"github.com/lib/pq"
)

const (
	DefaultStakerRewardsUpdateInterval = 7
	DefaultCommunityPoolUpdateInterval = 5
)

// DistributionParamsRow represents a single row inside the distribution_params table
type DistributionParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}

// -------------------------------------------------------------------------------------------------------------------

// CommunityPoolRow represents a single row inside the total_supply table
type CommunityPoolRow struct {
	OneRowID bool        `db:"one_row_id"`
	Coins    *DbDecCoins `db:"coins"`
	Height   int64       `db:"height"`
}

// NewCommunityPoolRow allows to easily create a new CommunityPoolRow
func NewCommunityPoolRow(coins DbDecCoins, height int64) CommunityPoolRow {
	return CommunityPoolRow{
		OneRowID: true,
		Coins:    &coins,
		Height:   height,
	}
}

// Equals return true if one CommunityPoolRow representing the same row as the original one
func (v CommunityPoolRow) Equals(w CommunityPoolRow) bool {
	return v.Coins.Equal(w.Coins) &&
		v.Height == w.Height
}

// DbOperatorRewardProportion represents a single operator reward proportion record stored in the database.
// The RewardProportion is stored as a string representation of cosmos.Dec.
type DbOperatorRewardProportion struct {
	OperatorAddr     string
	RewardProportion string // cosmos.Dec as string
}

// Value implements the driver.Valuer interface to convert DbOperatorRewardProportion
// into a Postgres composite type string format "(operator_addr,reward_proportion)".
func (d *DbOperatorRewardProportion) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", d.OperatorAddr, d.RewardProportion), nil
}

// Scan implements the sql.Scanner interface to parse a Postgres composite type string
// into a DbOperatorRewardProportion struct.
func (d *DbOperatorRewardProportion) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("DbOperatorRewardProportion: expected []byte, got %T", src)
	}

	s := string(b)
	// Trim quotes and parentheses
	s = strings.Trim(s, `"`)
	s = strings.Trim(s, "()")

	// Split the composite string by comma into two parts
	parts := strings.SplitN(s, ",", 2)
	if len(parts) != 2 {
		return fmt.Errorf("DbOperatorRewardProportion: invalid format %q", s)
	}

	d.OperatorAddr = parts[0]
	d.RewardProportion = parts[1]
	return nil
}

// DbOperatorRewardProportions represents a slice of DbOperatorRewardProportion pointers.
type DbOperatorRewardProportions []*DbOperatorRewardProportion

// Value implements the driver.Valuer interface for DbOperatorRewardProportions.
// It converts the slice into a Postgres array of composite types.
func (list DbOperatorRewardProportions) Value() (driver.Value, error) {
	values := make([]string, len(list))
	for i, elem := range list {
		val, err := elem.Value()
		if err != nil {
			return nil, err
		}
		values[i] = val.(string)
	}
	// Use pq.Array to format as Postgres array
	return pq.Array(values).Value()
}

// Scan implements the sql.Scanner interface for DbOperatorRewardProportions.
// It scans a Postgres array of composite types into the slice.
func (list *DbOperatorRewardProportions) Scan(src interface{}) error {
	var raw []string
	// Use pq.Array to scan the Postgres array into a slice of strings
	err := pq.Array(&raw).Scan(src)
	if err != nil {
		return err
	}

	result := make(DbOperatorRewardProportions, len(raw))
	for i, r := range raw {
		var d DbOperatorRewardProportion
		// Scan each composite string into the struct
		err := d.Scan([]byte(r))
		if err != nil {
			return err
		}
		result[i] = &d
	}
	*list = result
	return nil
}

func NewDbOperatorRewardProportions(props []distrtypes.OperatorRewardProportion) DbOperatorRewardProportions {
	result := make([]*DbOperatorRewardProportion, len(props))
	for i, p := range props {
		result[i] = &DbOperatorRewardProportion{
			OperatorAddr:     p.OperatorAddr,
			RewardProportion: p.RewardProportion.String(),
		}
	}
	return result
}

type DistributionIndexerParams struct {
	CommunityPoolUpdateInterval *int64
	UpdateTime                  *time.Time
}
