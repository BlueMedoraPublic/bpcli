package cmd

import (
	"github.com/spf13/cobra"
)

var logAgentTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage log agent tasks",
}

func init() {
	logAgentCmd.AddCommand(logAgentTaskCmd)
}
