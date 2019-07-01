package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var listCredentialCmd = &cobra.Command{
	Use:   "list",
	Short: "List configured credentials",
	Run: func(cmd *cobra.Command, args []string) {
		listCredentials()
	},
}

func init() {
	credentialCmd.AddCommand(listCredentialCmd)
	listCredentialCmd.Flags().StringVar(&credentialID, "id", "", "The ID of the SourceTemplate")
}

func listCredentials() {
	x, err := bp.GetCredentials()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, credential := range x {
		if err := credential.Print(false); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
