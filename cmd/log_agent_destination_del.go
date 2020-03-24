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
}

func delDestinationsLogAgent() error {
	return bp.DeleteLogAgentDest(logAgentID, logConfigID)
}
