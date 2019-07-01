package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var getCredentialTypeCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns related configuration for a credential type",
	Run: func(cmd *cobra.Command, args []string) {
		getcredentialTypebyID()
	},
}

func init() {
	credentialTypeCmd.AddCommand(getCredentialTypeCmd)
	getCredentialTypeCmd.Flags().StringVar(&credentialTypeID, "id", "", "The ID of the SourceTemplate")
	getCredentialTypeCmd.MarkFlagRequired("id")
}

func getcredentialTypebyID() {
	x, err := bp.GetCredentialType(credentialTypeID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	x.Print()
}
