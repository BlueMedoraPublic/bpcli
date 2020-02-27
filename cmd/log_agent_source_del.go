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
	logAgentSourceDelete.Flags().StringVar(&logAgentID, "agent-id", "", "The ID of the log agent")
	logAgentSourceDelete.Flags().StringVar(&logAgentSourceID, "source-id", "", "The ID of the log agent's source")
	logAgentSourceDelete.MarkFlagRequired("agent-id")
	logAgentSourceDelete.MarkFlagRequired("source-id")
}

func delSourceLogAgent() error {
	return bp.DeleteLogAgentSource(logAgentID, logAgentSourceID)
}
