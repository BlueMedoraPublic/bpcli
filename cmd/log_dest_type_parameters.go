package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logDestTypeParametersCmd = &cobra.Command{
	Use:   "parameters",
	Short: "Describe a destination type's parameters, for easy templating",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogDestTypeParameters(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestTypeCmd.AddCommand(logDestTypeParametersCmd)
	logDestTypeParametersCmd.Flags().StringVar(&logDestTypeID, "id", "", "The ID of the destination type")
	logDestTypeParametersCmd.MarkFlagRequired("id")
}

func getLogDestTypeParameters() error {
	return nil
}
