package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var delCredentialCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a credential",
	Run: func(cmd *cobra.Command, args []string) {
		delCredential()
	},
}

func init() {
	credentialCmd.AddCommand(delCredentialCmd)
	delCredentialCmd.Flags().StringVar(&credentialID, "id", "", "The ID of the SourceTemplate")
	delCredentialCmd.MarkFlagRequired("id")
}

func delCredential() {
	err := bp.DeleteCredential(credentialID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("credential deleted:", credentialID)
}
