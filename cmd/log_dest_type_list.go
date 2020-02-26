package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logDestTypeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all destination types",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listLogDestTypes(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestTypeCmd.AddCommand(logDestTypeListCmd)
}

func listLogDestTypes() error {
	return nil
}
