package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentInstallCmd = &cobra.Command{
	Use:   "install-cmd",
	Short: "Get the command(s) required to install an agent on an OS",
	Run: func(cmd *cobra.Command, args []string) {
		if err := agentInstallCommand(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentCmd.AddCommand(logAgentInstallCmd)
}

func agentInstallCommand() error {
	b, err := bp.InstallCMDLogAgent()
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
