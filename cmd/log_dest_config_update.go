package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logDestConfigUpdateVersion = &cobra.Command{
	Use:   "update-version",
	Short: "Update an existing configuration's version to the latest",
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateVersionLogDestConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestConfigCmd.AddCommand(logDestConfigUpdateVersion)
	logDestConfigUpdateVersion.Flags().StringVar(&logDestConfigID, "id", "", "The ID of the destination config")
	logDestConfigUpdateVersion.MarkFlagRequired("id")
}

func updateVersionLogDestConfigs() error {
	return nil
}
