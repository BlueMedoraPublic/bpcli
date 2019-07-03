package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCollectorCmd = &cobra.Command{
	Use:   "list",
	Short: "List collectors",
	Run: func(cmd *cobra.Command, args []string) {
		listCollector()
	},
}

func init() {
	collectorCmd.AddCommand(listCollectorCmd)
}

func listCollector() {
	c, err := bp.GetCollectors()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, group := range c {
		for _, collector := range group.Collectors {
			if err := collector.Print(jsonFmt); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		}

	}
}
