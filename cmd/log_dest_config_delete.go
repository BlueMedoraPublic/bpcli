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
	logDestConfigDeleteCmd.Flags().StringVar(&logDestConfigID, "id", "", "The ID of the destination config")
	logDestConfigDeleteCmd.MarkFlagRequired("id")
}

func delLogDestConfigs() error {
	return bp.DelLogDestConfig(logDestConfigID)
}
