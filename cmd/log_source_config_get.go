package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logSourceConfigGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Describe a source config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogSourceConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsConfigCmd.AddCommand(logSourceConfigGetCmd)
}

func getLogSourceConfigs() error {
	c, err := bp.GetLogSourceConfig(logConfigID)
	if err != nil {
		return err
	}

	return c.Print(jsonFmt)
}
