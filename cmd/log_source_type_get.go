package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logSourceTypeGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Describe a source type",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogSourceTypes(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsTypeCmd.AddCommand(logSourceTypeGetCmd)
}

func getLogSourceTypes() error {
	s, err := bp.GetLogSourceType(logConfigID)
	if err != nil {
		return err
	}
	return s.Print(jsonFmt)
}
