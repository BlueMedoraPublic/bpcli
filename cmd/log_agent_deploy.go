package cmd

import (
	"github.com/spf13/cobra"
)

var logAgentDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy source and destination configuration to an agent",
}

func init() {
	logAgentCmd.AddCommand(logAgentDeployCmd)
}
