package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logDestConfigDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a destination config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := delLogDestConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestConfigCmd.AddCommand(logDestConfigDeleteCmd)
}

func delLogDestConfigs() error {
	return bp.DelLogDestConfig(logConfigID)
}
