package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var getcredentialTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Returns credential  configured credentials of a specific type",
	Run: func(cmd *cobra.Command, args []string) {
		getcredentialTemplateByID()
	},
}

func init() {
	credentialTypeCmd.AddCommand(getcredentialTemplateCmd)
	getcredentialTemplateCmd.Flags().StringVar(&credentialTypeID, "id", "", "The ID of the SourceTemplate")
	getcredentialTemplateCmd.MarkFlagRequired("id")
}

func getcredentialTemplateByID() {
	c, err := bp.GetCredentialTemplate(credentialTypeID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	c.Print()
}
