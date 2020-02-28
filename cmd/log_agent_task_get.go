package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentTaskGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Describe a log agent task",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogAgentTask(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentTaskCmd.AddCommand(logAgentTaskGetCmd)
}

func getLogAgentTask() error {
	a, err := bp.GetLogAgentTask(logAgentID, logTaskID)
	if err != nil {
		return err
	}

	return a.Print(jsonFmt)
}
