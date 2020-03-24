package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logSourceConfigDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a source config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := delLogSourceConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsConfigCmd.AddCommand(logSourceConfigDeleteCmd)
}

func delLogSourceConfigs() error {
	if err := bp.DeleteLogSourceConfig(logConfigID); err != nil {
		return err
	}

	fmt.Println("log source config " + logConfigID + " deleted")
	return nil
}
