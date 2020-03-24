package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logSourceTypeParametersCmd = &cobra.Command{
	Use:   "parameters",
	Short: "Describe a source types parameters, for easy templating",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogSourceTypeParameters(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsTypeCmd.AddCommand(logSourceTypeParametersCmd)
}

func getLogSourceTypeParameters() error {
	s, err := bp.GetLogSourceTypeParameters(logConfigID)
	if err != nil {
		return err
	}
	fmt.Println(string(s))
	return nil
}
