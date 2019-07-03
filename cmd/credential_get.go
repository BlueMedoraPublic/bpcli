package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var getCredentialCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a credential",
	Run: func(cmd *cobra.Command, args []string) {
		getCredentialByID()
	},
}

func init() {
	credentialCmd.AddCommand(getCredentialCmd)
	getCredentialCmd.Flags().StringVar(&credentialID, "id", "", "The ID of the SourceTemplate")
	getCredentialCmd.MarkFlagRequired("id")
}

func getCredentialByID() {
	x, err := bp.GetCredential(credentialID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := x.Print(jsonFmt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
