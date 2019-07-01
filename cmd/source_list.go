package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var listSourceCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured sources",
	Run: func(cmd *cobra.Command, args []string) {
		listSources()
	},
}

func init() {
	sourceCmd.AddCommand(listSourceCmd)
}

func listSources() {
	s, err := bp.GetSources()
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
