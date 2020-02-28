package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentSourceGet = &cobra.Command{
	Use:   "get",
	Short: "Describe a source configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getSourceLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentSourceCmd.AddCommand(logAgentSourceGet)
}

func getSourceLogAgent() error {
	s, err := bp.GetLogAgentSource(logAgentID, logConfigID)
	if err != nil {
		return err
	}

	return s.Print(jsonFmt)
}
