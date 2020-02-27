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
	logDestConfigGetCmd.Flags().StringVar(&logDestConfigID, "id", "", "The ID of the destination config")
	logDestConfigGetCmd.MarkFlagRequired("id")
}

func getLogDestConfigs() error {
	d, err := bp.GetLogDestConfig(logDestConfigID)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
