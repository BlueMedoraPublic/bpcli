package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var sourceTypeTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Returns source template for a given source type",
	Run: func(cmd *cobra.Command, args []string) {
		sourceTypeTemplateCmdByID()
	},
}

func init() {
	sourceTypeCmd.AddCommand(sourceTypeTemplateCmd)
	sourceTypeTemplateCmd.Flags().StringVar(&sourceTemplateID, "id", "", "The ID of the SourceTemplate")
	sourceTypeTemplateCmd.MarkFlagRequired("id")
}

func sourceTypeTemplateCmdByID() {
	s, err := bp.GetSourceTemplate(sourceTemplateID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err := s.Print(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
