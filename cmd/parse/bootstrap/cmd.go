package bootstrap

import (
	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewBootstrapCmd returns the Cobra command that allows to update states for bootstrap
func NewBootstrapCmd(parseCfg *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bootstrap",
		Short: "update states for bootstrap",
	}

	cmd.AddCommand(
		bootstrapCmd(parseCfg),
	)

	return cmd
}
