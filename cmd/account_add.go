package cmd

import (
	"github.com/spf13/cobra"
)

var addAccountCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a BindPlane account to the list",
}

func init() {
	accountCmd.AddCommand(addAccountCmd)
	addAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
	addAccountCmd.Flags().StringVar(&accountID, "id", "", "The BindPlane API Key")
}
