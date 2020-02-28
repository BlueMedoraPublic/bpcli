package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentSourceDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a source configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if err := delSourceLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentSourceCmd.AddCommand(logAgentSourceDelete)
}

func delSourceLogAgent() error {
	return bp.DeleteLogAgentSource(logAgentID, logConfigID)
}
