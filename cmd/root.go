package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/bpcli/bindplane/sdk"
	"github.com/BlueMedoraPublic/bpcli/util/uuid"

	"github.com/spf13/cobra"
)

// flags
var accountName string
var sourceFile string
var credentialFile string
var jsonFmt bool
var zshCompletion bool
var watch bool

// uuid flags
var accountKey string
var jobID string
var groupID string
var collectorID string
var sourceTypeID string
var sourceID string
var credentialID string
var credentialTypeID string
var sourceTemplateID string
var logSourceTypeID string
var logSourceConfigID string
var logAgentID string
var logTaskID string
var logAgentSourceID string
var logAgentDestID string
var logDestTypeID string
var logDestConfigID string

// bindplane object
var bp sdk.BindPlane

var rootCmd = &cobra.Command{
	Use:   "bpcli",
	Short: "Command line utility for interacting with the BindPlane API",
}

// Execute is called by the cobra framework
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&jsonFmt, "json", false, "enable json output for commands that support json")
}

func initConfig() {
	// avoid running bp.Init() if these commands were passed
	// as argument one
	y := []string{"help", "version", "completion", "account"}
	for _, subCmd := range y {
		if subCmd == os.Args[1] {
			return
		}
	}

	// Init the configuration and set defaults
	err := bp.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := validateFlags(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

}

// validateFlags returns nil if all flags pass their checks
func validateFlags() error {
	if len(collectorID) > 0 {
		if uuid.IsUUID(collectorID) != true {
			return errors.New("collector id must be a valid UUID")
		}
	}

	if len(groupID) > 0 {
		if uuid.IsUUID(groupID) != true {
			return errors.New("group id must be a valid UUID")
		}
	}

	if len(jobID) > 0 {
		if uuid.IsUUID(jobID) != true {
			return errors.New("job id must be a valid UUID")
		}
	}

	if len(sourceID) > 0 {
		if uuid.IsUUID(sourceID) != true {
			return errors.New("source id must be a valid UUID")
		}
	}

	if len(credentialID) > 0 {
		if uuid.IsUUID(credentialID) != true {
			return errors.New("credential must be a valid UUID")
		}
	}

	if len(credentialTypeID) > 0 {
		if uuid.IsUUID(credentialTypeID) != true {
			return errors.New("credential type id must be a valid UUID")
		}
	}

	return nil
}
