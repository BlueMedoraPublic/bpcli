package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCollectorGroupCmd = &cobra.Command{
	Use:   "list",
	Short: "List collector groups",
	Run: func(cmd *cobra.Command, args []string) {
		listCollectorGroups()
	},
}

func init() {
	collectorGroupCmd.AddCommand(listCollectorGroupCmd)
}

func listCollectorGroups() {
	c, err := bp.GetCollectors()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, group := range c {
		if err := group.Print(jsonFmt); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	}
}
