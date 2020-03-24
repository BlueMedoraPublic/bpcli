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
	a, err := bp.ListLogAgents()
	if err != nil {
		return err
	}

	for _, agent := range a {
		if err := agent.Print(jsonFmt); err != nil {
			return err
		}
	}
	return nil
}
