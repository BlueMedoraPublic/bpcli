package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logSourceConfigUpdateVersion = &cobra.Command{
	Use:   "update-version",
	Short: "Update an existing configuration's version to the latest",
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateVersionLogSourceConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsConfigCmd.AddCommand(logSourceConfigUpdateVersion)
}

func updateVersionLogSourceConfigs() error {
	c, err := bp.UpdateVersionLogSourceConfig(logConfigID)
	if err != nil {
		return err
	}

	return c.Print(jsonFmt)
}
