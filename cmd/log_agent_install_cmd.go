package cmd

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/pkg/errors"
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

	p := []string{}
	for platform, command := range platforms {
		p = append(p, platform)
		if platform == logAgentPlatform {
			fmt.Println(command)
			return nil
		}
	}

	// exit early if no platforms are returned
	if len(p) < 1 {
		return errors.New("unexpected response from server, no install commands returned")
	}

	// safe to index p[0] because we know the slice is at
	// least length 1 from the check above ^
	foundPlatforms := p[0]
	for i, p := range p {
		if i == 1 {
			continue
		}
		foundPlatforms = foundPlatforms + ", " + p
	}

	err = errors.New("platform is not supported: " + logAgentPlatform)
	return errors.Wrap(err, "supported platforms: " + foundPlatforms)
}
