package cmd

import (
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/bpcli/config"
	"github.com/spf13/cobra"
)

var removeAccountCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a BindPlane account",
	Run: func(cmd *cobra.Command, args []string) {
		remove()
	},
}

func init() {
	accountCmd.AddCommand(removeAccountCmd)
	removeAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
	removeAccountCmd.MarkFlagRequired("name")
}

func remove() {
	if err := config.Remove(accountName); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(accountName + " has been successfully removed")
}
