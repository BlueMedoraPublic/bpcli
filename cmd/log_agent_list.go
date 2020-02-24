package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all log agents",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listLogAgents(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentCmd.AddCommand(logAgentListCmd)
}

func listLogAgents() error {
	return nil
}
