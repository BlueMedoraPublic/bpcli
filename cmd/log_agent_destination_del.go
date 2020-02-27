package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDestDel = &cobra.Command{
	Use:   "delete",
	Short: "Delete a destination configurations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := delDestinationsLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentDestinationCmd.AddCommand(logAgentDestDel)
	logAgentDestDel.Flags().StringVar(&logAgentID, "agent-id", "", "The ID of the log agent")
	logAgentDestDel.Flags().StringVar(&logDestConfigID, "destination-id", "", "The destination configuration ID")
	logAgentDestDel.MarkFlagRequired("agent-id")
	logAgentDestDel.MarkFlagRequired("destination-id")
}

func delDestinationsLogAgent() error {
	return bp.DeleteLogAgentDest(logAgentID, logDestConfigID)
}
