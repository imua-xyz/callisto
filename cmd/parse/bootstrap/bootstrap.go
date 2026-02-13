package bootstrap

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/modules"
	modsregistrar "github.com/forbole/juno/v5/modules/registrar"
	"github.com/forbole/juno/v5/parser"
	"github.com/forbole/juno/v5/types/config"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
)

func getParserContext(cfg config.Config, parseConfig *parsecmdtypes.Config) (*parser.Context, error) {
	// Build the codec
	encodingConfig := parseConfig.GetEncodingConfigBuilder()()

	// Get the db
	databaseCtx := database.NewContext(cfg.Database, encodingConfig, parseConfig.GetLogger())
	db, err := parseConfig.GetDBBuilder()(databaseCtx)
	if err != nil {
		return nil, err
	}

	// Setup the logging
	err = parseConfig.GetLogger().SetLogFormat(cfg.Logging.LogFormat)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = parseConfig.GetLogger().SetLogLevel(cfg.Logging.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	// Get the modules
	context := modsregistrar.NewContext(cfg, nil, encodingConfig, db, nil, parseConfig.GetLogger())
	mods := parseConfig.GetRegistrar().BuildModules(context)
	registeredModules := modsregistrar.GetModules(mods, cfg.Chain.Modules, parseConfig.GetLogger())

	return parser.NewContext(encodingConfig, nil, db, parseConfig.GetLogger(), registeredModules), nil
}

// startParsing represents the function that should be called when the parse command is executed
func startParsing(ctx *parser.Context) error {
	// Start periodic operations
	scheduler := gocron.NewScheduler(time.UTC)
	for _, module := range ctx.Modules {
		if module, ok := module.(modules.PeriodicOperationsModule); ok {
			err := module.RegisterPeriodicOperations(scheduler)
			if err != nil {
				return err
			}
		}
	}
	scheduler.StartAsync()

	// Run all the async operations
	for _, module := range ctx.Modules {
		if module, ok := module.(modules.AsyncOperationsModule); ok {
			go module.RunAsyncOperations()
		}
	}

	// Listen for shutdown signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigCh
	ctx.Logger.Info("caught signal; shutting down...", "signal", sig.String())

	// close the database.
	ctx.Database.Close()
	return nil
}

// bootstrapCmd returns a Cobra command that allows to update states for bootstrap
func bootstrapCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Update the bootstrap states to indexer automatically",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := getParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}
			// Run all the additional operations
			for _, module := range parseCtx.Modules {
				if module, ok := module.(modules.AdditionalOperationsModule); ok {
					err = module.RunAdditionalOperations()
					if err != nil {
						return err
					}
				}
			}

			return startParsing(parseCtx)
		},
	}
}
