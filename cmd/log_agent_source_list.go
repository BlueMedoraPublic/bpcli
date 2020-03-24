package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentSourceList = &cobra.Command{
	Use:   "list",
	Short: "List source configurations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listSourceLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentSourceCmd.AddCommand(logAgentSourceList)
}

func listSourceLogAgent() error {
	s, err := bp.ListLogAgentSources(logAgentID)
	if err != nil {
		return err
	}

	for _, source := range s {
		if err := source.Print(jsonFmt); err != nil {
			return err
		}
	}
	return nil
}
