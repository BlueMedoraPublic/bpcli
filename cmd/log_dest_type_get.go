package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logDestTypeGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Describe a destination type",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogDestTypes(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestTypeCmd.AddCommand(logDestTypeGetCmd)
}

func getLogDestTypes() error {
	d, err := bp.GetLogDestType(logConfigID)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
