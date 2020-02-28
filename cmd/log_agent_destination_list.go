package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logAgentDestList = &cobra.Command{
	Use:   "list",
	Short: "List destination configurations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listDestinationsLogAgent(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logAgentDestinationCmd.AddCommand(logAgentDestList)
}

func listDestinationsLogAgent() error {
	d, err := bp.ListLogAgentDest(logAgentID)
	if err != nil {
		return err
	}

	for _, d := range d {
		if err := d.Print(jsonFmt); err != nil {
			return err
		}
	}
	return nil
}
