package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDeploySource = &cobra.Command{
	Use:   "source",
	Short: "Deploy a source configuration to a log agent",
	Run: func(cmd *cobra.Command, args []string) {
		if err := deploySourceLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentDeployCmd.AddCommand(logAgentDeploySource)
	logAgentDeploySource.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentDeploySource.Flags().StringVar(&sourceFile, "file", "", "The source json file")
	logAgentDeploySource.MarkFlagRequired("id")
}

func deploySourceLogAgent() error {
	return nil
}
