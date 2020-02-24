package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentUpdateVersion = &cobra.Command{
	Use:   "update-version",
	Short: "Update a log agent's version to the latest",
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateVersionLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentCmd.AddCommand(logAgentUpdateVersion)
	logAgentUpdateVersion.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentUpdateVersion.MarkFlagRequired("id")
}

func updateVersionLogAgent() error {
	return nil
}
