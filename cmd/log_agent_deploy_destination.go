package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDeployDest = &cobra.Command{
	Use:   "destinaton",
	Short: "Deploy a destination configuration to a log agent",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deployDestinationLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentDeployCmd.AddCommand(logAgentDeployDest)
	logAgentDeployDest.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentDeployDest.Flags().StringVar(&sourceFile, "file", "", "The source json file")
	logAgentDeployDest.MarkFlagRequired("id")
}

func deployDestinationLogAgent() error {
	return nil
}
