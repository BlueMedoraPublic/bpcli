package cmd

import (
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/bpcli/config"
	"github.com/spf13/cobra"
)

var addAccountCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a BindPlane account",
	Run: func(cmd *cobra.Command, args []string) {
		add()
	},
}

func init() {
	accountCmd.AddCommand(addAccountCmd)
	addAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
	addAccountCmd.Flags().StringVar(&accountKey, "key", "", "The BindPlane API Key")
	addAccountCmd.MarkFlagRequired("name")
	addAccountCmd.MarkFlagRequired("key")
}

func add() {
	if err := config.AddAccount(accountName, accountKey); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(accountName + " has been successfully added")
}
