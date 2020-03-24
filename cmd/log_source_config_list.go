package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logSourceConfigListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all source configs",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listLogSourceConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsConfigCmd.AddCommand(logSourceConfigListCmd)
}

func listLogSourceConfigs() error {
	c, err := bp.ListLogSourceConfigs()
	if err != nil {
		return err
	}

	for _, config := range c {
		if err := config.Print(jsonFmt); err != nil {
			return err
		}
	}
	return nil
}
