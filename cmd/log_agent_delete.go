package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a log agent",
	Run: func(cmd *cobra.Command, args []string) {
		if err := delLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentCmd.AddCommand(logAgentDeleteCmd)
	logAgentDeleteCmd.Flags().StringVar(&logAgentID, "id", "", "The ID of the log agent")
	logAgentDeleteCmd.MarkFlagRequired("id")
}

func delLogAgent() error {
	if err := bp.DeleteLogAgent(logAgentID); err != nil {
		return err
	}

	fmt.Println("log agent " + logAgentID + " deleted")
	return nil
}
