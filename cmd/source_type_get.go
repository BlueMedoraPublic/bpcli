package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var sourceTypeGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a source type description for a given id",
	Run: func(cmd *cobra.Command, args []string) {
		getSourceType()
	},
}

func init() {
	sourceTypeCmd.AddCommand(sourceTypeGetCmd)
	sourceTypeGetCmd.Flags().StringVar(&sourceTypeID, "id", "", "The ID of the SourceType")
	sourceTypeGetCmd.MarkFlagRequired("id")
}

func getSourceType() {
	s, err := bp.GetSourceType(sourceTypeID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := s.Print(jsonFmt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
