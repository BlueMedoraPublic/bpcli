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
}

func deployDestinationLogAgent() error {
	d, err := bp.DeployLogAgentDest(logAgentID, logConfigID)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
