package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentSourceDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a source configuration to a log agent",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deploySourceLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentSourceCmd.AddCommand(logAgentSourceDeploy)
	logAgentSourceDeploy.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentSourceDeploy.Flags().StringVar(&sourceFile, "file", "", "The source json file")
	logAgentSourceDeploy.MarkFlagRequired("id")
	logAgentSourceDeploy.MarkFlagRequired("file")
}

func deploySourceLogAgent() error {
	return nil
}
