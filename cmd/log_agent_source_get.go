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
	logAgentSourceGet.Flags().StringVar(&logAgentID, "agent-id", "", "The ID of the log agent")
	logAgentSourceGet.Flags().StringVar(&logAgentSourceID, "source-id", "", "The ID of the log agent's source")
	logAgentSourceGet.MarkFlagRequired("agent-id")
	logAgentSourceGet.MarkFlagRequired("source-id")
}

func getSourceLogAgent() error {
	s, err := bp.GetLogAgentSource(logAgentID, logAgentSourceID)
	if err != nil {
		return err
	}

	return s.Print(jsonFmt)
}
