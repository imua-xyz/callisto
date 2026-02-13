package oracle

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/x/authz"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/callisto/v4/types"
	juno "github.com/forbole/juno/v5/types"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"
)

// HandleMsgExec implements AuthzMessageModule. It handles the case wherein
// a grantee is permitted to execute a message on behalf of the granter.
func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
	return m.HandleMsg(index, executedMsg, tx)
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	if _, ok := msg.(*oracletypes.MsgCreatePrice); ok {
		return m.handlePriceEvents(tx)
	}
	// we do not handle MsgUpdateParams because it is handled in HandleBlock
	return nil
}

// handlePriceEvents handles the creation of price events
func (m *Module) handlePriceEvents(tx *juno.Tx) error {
	events := juno.FindEventsByType(tx.Events, oracletypes.EventTypeCreatePrice)
	for _, event := range events {
		finalPrice, err := juno.FindAttributeByKey(event, oracletypes.AttributeKeyFinalPrice)
		if err != nil {
			// the x/oracle module reuses the event name across different types
			// so, if this attribute is not found, it means that this event is not
			// relevant to us
			continue
		}
		split := strings.Split(finalPrice.Value, "_")
		if len(split) != 4 {
			return fmt.Errorf("invalid price event: %s", finalPrice.Value)
		}
		price := split[2]
		// check if price is a valid float
		_, err = strconv.ParseFloat(price, 64)
		if err != nil {
			// base64 encoded price applies only to NSTs, and will fail
			// ParseFloat, so we can safely ignore these. NST handling
			// is performed in x/assets and x/delegation separately.
			continue
		}
		tokenID := split[0]
		roundID := split[1]
		priceDecimals := split[3]
		timestamp, err := time.Parse(time.RFC3339, tx.Timestamp)
		if err != nil {
			return fmt.Errorf("error parsing timestamp: %w", err)
		}
		priceHistory := types.NewOraclePriceHistoryFromStr(
			tokenID, roundID,
			price, priceDecimals,
			timestamp.Format(oracletypes.TimeLayout),
		)
		if err := m.db.SaveOraclePriceHistory(priceHistory); err != nil {
			return fmt.Errorf("error saving oracle price history: %w", err)
		}
		// this event converts the round to committable state
		// the endblocker then moves it to closed state and increases
		// the roundID. we do not handle that in endblocker because
		// it is difficult to look up from the emitted event.
		// instead, we handle it here.
		if err := m.db.IncreaseNextRoundID(tokenID); err != nil {
			return fmt.Errorf("error increasing next round ID: %w", err)
		}
	}
	return nil
}
