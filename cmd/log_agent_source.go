package cmd

import (
	"github.com/spf13/cobra"
)

var logAgentSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage log sources",
}

func init() {
	logAgentCmd.AddCommand(logAgentSourceCmd)
}
