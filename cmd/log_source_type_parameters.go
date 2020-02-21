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
	logSourceTypeParametersCmd.Flags().StringVar(&logSourceTypeID, "id", "", "The ID of the source type")
	logSourceTypeParametersCmd.MarkFlagRequired("id")
}

func getLogSourceTypeParameters() error {
	s, err := bp.GetLogSourceTypeParameters(logSourceTypeID)
	if err != nil {
		return err
	}
	fmt.Println(string(s))
	//return s.Print(jsonFmt)
	return nil
}
