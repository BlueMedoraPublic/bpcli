package cmd

import (
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/bpcli/config"
	"github.com/spf13/cobra"
)

var addAccountCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a BindPlane account to the list",
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

func init() {
	accountCmd.AddCommand(addAccountCmd)
	addAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
	addAccountCmd.Flags().StringVar(&accountID, "id", "", "The BindPlane API Key")
	setAccountCmd.MarkFlagRequired("name")
	setAccountCmd.MarkFlagRequired("id")
}

func add() {
	err := config.AddAccount(accountName, accountID)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
