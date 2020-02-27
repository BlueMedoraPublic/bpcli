package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDestGet = &cobra.Command{
	Use:   "get",
	Short: "Describe a destination configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getDestinationsLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentDestinationCmd.AddCommand(logAgentDestGet)
	logAgentDestGet.Flags().StringVar(&logAgentID, "agent-id", "", "The ID of the log agent")
	logAgentDestGet.Flags().StringVar(&logDestConfigID, "destination-id", "", "The destination configuration ID")
	logAgentDestGet.MarkFlagRequired("agent-id")
	logAgentDestGet.MarkFlagRequired("destination-id")
}

func getDestinationsLogAgent() error {
	d, err := bp.GetLogAgentDest(logAgentID, logDestConfigID)
	if err != nil {
		return err
	}

 	return d.Print(jsonFmt)
}
