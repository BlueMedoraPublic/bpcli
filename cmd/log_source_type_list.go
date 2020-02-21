package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logSourceTypeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all source types",
	Run: func(cmd *cobra.Command, args []string) {
		listLogSourceTypes()
	},
}

func init() {
	logsTypeCmd.AddCommand(logSourceTypeListCmd)
}

func listLogSourceTypes() {
	s, err := bp.ListLogSourceTypes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, sourceType := range s {
		if err := sourceType.Print(false); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
