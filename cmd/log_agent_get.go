package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Describe a log agent",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentCmd.AddCommand(logAgentGetCmd)
	logAgentGetCmd.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentGetCmd.MarkFlagRequired("id")
}

func getLogAgent() error {
	a, err := bp.GetLogAgent(logAgentID)
	if err != nil {
		return err
	}

	return a.Print(jsonFmt)
}
