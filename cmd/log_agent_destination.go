package cmd

import (
	"github.com/spf13/cobra"
)

var logAgentDestinationCmd = &cobra.Command{
	Use:   "destination",
	Short: "Manage log destinations",
}

func init() {
	logAgentCmd.AddCommand(logAgentDestinationCmd)
}
