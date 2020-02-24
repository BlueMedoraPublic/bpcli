package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDestDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a destination configuration to a log agent",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deployDestinationLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentDestinationCmd.AddCommand(logAgentDestDeploy)
	logAgentDestDeploy.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentDestDeploy.Flags().StringVar(&sourceFile, "file", "", "The source json file")
	logAgentDestDeploy.MarkFlagRequired("id")
	logAgentDestDeploy.MarkFlagRequired("file")
}

func deployDestinationLogAgent() error {
	return nil
}
