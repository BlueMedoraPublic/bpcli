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
	logAgentSourceList.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentSourceList.MarkFlagRequired("id")
}

func listSourceLogAgent() error {
	return nil
}
