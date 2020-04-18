package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/bpcli/util/pprint"

	"github.com/spf13/cobra"
)

var logAgentInstallCmd = &cobra.Command{
	Use:   "install-cmd",
	Short: "Get the command required to install an agent, with an optional template",
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
	c, err := bp.InstallCMDLogAgent(logAgentPlatform, logTemplateID)
	if err != nil {
		return err
	}

	if logAgentPlatform == "all" {
		m := make(map[string]string)
		if err := json.Unmarshal([]byte(c), &m); err != nil {
			return err
		}
		return pprint.PrintJSONStringMap(m)
	}

	fmt.Println(c)
	return nil
}
