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
		if err := listLogSourceTypes(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsTypeCmd.AddCommand(logSourceTypeListCmd)
}

func listLogSourceTypes() error {
	s, err := bp.ListLogSourceTypes()
	if err != nil {
		return err
	}

	for _, sourceType := range s {
		if err := sourceType.Print(false); err != nil {
			return err
		}
	}
	return nil
}
