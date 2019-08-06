package cmd

import (
	"github.com/spf13/cobra"
)

var removeAccountCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a BindPlane account from the list",
}

func init() {
	accountCmd.AddCommand(removeAccountCmd)
	removeAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
}
