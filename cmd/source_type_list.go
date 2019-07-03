package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sourceTypeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all source types",
	Run: func(cmd *cobra.Command, args []string) {
		listSourceTypes()
	},
}

func init() {
	sourceTypeCmd.AddCommand(sourceTypeListCmd)
}

func listSourceTypes() {
	s, err := bp.ListSourceTypes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, source := range s {
		if err := source.Print(false); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
