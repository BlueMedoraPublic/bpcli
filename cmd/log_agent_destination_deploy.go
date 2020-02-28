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
	logAgentDestDeploy.Flags().StringVar(&logAgentID, "agent-id", "", "The ID of the log agent")
	logAgentDestDeploy.Flags().StringVar(&logAgentDestID, "config-id", "", "The destination config id")
	logAgentDestDeploy.MarkFlagRequired("agent-id")
	logAgentDestDeploy.MarkFlagRequired("config-id")
}

func deployDestinationLogAgent() error {
	d, err := bp.DeployLogAgentDest(logAgentID, logAgentDestID)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
