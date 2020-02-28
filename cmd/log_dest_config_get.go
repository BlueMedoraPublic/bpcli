package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logDestConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Describe a destination config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogDestConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestConfigCmd.AddCommand(logDestConfigGetCmd)
}

func getLogDestConfigs() error {
	d, err := bp.GetLogDestConfig(logConfigID)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
