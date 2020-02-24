package cmd

import (
	"github.com/spf13/cobra"
)

var logAgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Manage BindPlane log agents",
}

func init() {
	logsCmd.AddCommand(logAgentCmd)
}
