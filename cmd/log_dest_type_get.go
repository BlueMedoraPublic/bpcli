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
	logDestTypeGetCmd.Flags().StringVar(&logDestTypeID, "id", "", "The ID of the destination type")
	logDestTypeGetCmd.MarkFlagRequired("id")
}

func getLogDestTypes() error {
	d, err := bp.GetLogDestType(logDestTypeID)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
