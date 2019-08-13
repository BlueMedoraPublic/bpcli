package cmd

import (
	"github.com/BlueMedoraPublic/bpcli/config"
	"github.com/spf13/cobra"
)

var removeAccountCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a BindPlane account from the list",
	Run: func(cmd *cobra.Command, args []string) {
		config.Remove(accountName)
	},
}

func init() {
	accountCmd.AddCommand(removeAccountCmd)
	removeAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
}
