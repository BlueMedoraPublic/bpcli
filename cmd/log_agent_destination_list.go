package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDestList = &cobra.Command{
	Use:   "list",
	Short: "List destination configurations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listDestinationsLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentDestinationCmd.AddCommand(logAgentDestList)
	logAgentDestList.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentDestList.MarkFlagRequired("id")
}

func listDestinationsLogAgent() error {
	return nil
}
