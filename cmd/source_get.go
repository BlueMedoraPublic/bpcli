package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var getSourceCmd = &cobra.Command{
	Use:   "get",
	Short: "Return a configured source",
	Run: func(cmd *cobra.Command, args []string) {
		getSourceByID()
	},
}

func init() {
	sourceCmd.AddCommand(getSourceCmd)
	getSourceCmd.Flags().StringVar(&sourceID, "id", "", "The ID of the source")
	getSourceCmd.MarkFlagRequired("id")
}

func getSourceByID() {
	s, err := bp.GetSource(sourceID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := s.Print(true); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
