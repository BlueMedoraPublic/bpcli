package cmd

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/spf13/cobra"
)

var logAgentInstallCmd = &cobra.Command{
	Use:   "install-cmd",
	Short: "Get the command required to install an agent",
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

	platforms := make(map[string]string)
	if err := json.Unmarshal(b, &platforms); err != nil {
		return err
	}

	for platform, command := range platforms {
		fmt.Println(platform, command)
	}
	return nil
}
