package cmd

import (
	"github.com/spf13/cobra"
)

var setAccountCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the BindPlane account to be used",
}

func init() {
	accountCmd.AddCommand(setAccountCmd)
	setAccountCmd.Flags().StringVar(&accountName, "name", "", "The name of the BindPlane account")
}
