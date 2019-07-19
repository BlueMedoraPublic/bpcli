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
var sourceFile string
var credentialFile string
var jsonFmt bool
var watch bool

// uuid flags
var jobID string
var groupID string
var collectorID string
var sourceTypeID string
var sourceID string
var credentialID string
var credentialTypeID string
var sourceTemplateID string

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
	y := []string{"help", "version"}
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
			return errors.New("--collectorID must be a valid UUID")
		}
	}

	if len(groupID) > 0 {
		if uuid.IsUUID(groupID) != true {
			return errors.New("--groupID must be a valid UUID")
		}
	}

	if len(jobID) > 0 {
		if uuid.IsUUID(jobID) != true {
			return errors.New("--jobID must be a valid UUID")
		}
	}

	if len(sourceID) > 0 {
		if uuid.IsUUID(sourceID) != true {
			return errors.New("--sourceID must be a valid UUID")
		}
	}

	if len(credentialID) > 0 {
		if uuid.IsUUID(credentialID) != true {
			return errors.New("--credentialID must be a valid UUID")
		}
	}

	if len(credentialTypeID) > 0 {
		if uuid.IsUUID(credentialTypeID) != true {
			return errors.New("--credentialTypeID must be a valid UUID")
		}
	}

	return nil
}
