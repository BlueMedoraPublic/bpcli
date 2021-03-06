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
}

func updateVersionLogAgent() error {
	t, err := bp.UpdateLogAgent(logAgentID)
	if err != nil {
		return err
	}

	return t.Print(jsonFmt)
}
