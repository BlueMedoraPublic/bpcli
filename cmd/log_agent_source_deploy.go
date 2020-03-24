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
}

func deploySourceLogAgent() error {
	d, err := bp.DeployLogAgentSource(logAgentID, logConfigID)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
