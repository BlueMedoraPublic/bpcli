package cmd

import (
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage BindPlane accounts",
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
