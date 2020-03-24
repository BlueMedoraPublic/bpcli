package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logDestConfigListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all destination configs",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listLogDestConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestConfigCmd.AddCommand(logDestConfigListCmd)
}

func listLogDestConfigs() error {
	d, err := bp.ListLogDestConfigs()
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
