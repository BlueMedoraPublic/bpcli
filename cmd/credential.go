package cmd

import (
	"github.com/spf13/cobra"
)

var credentialCmd = &cobra.Command{
	Use:   "credential",
	Short: "Manage credentials",
}

func init() {
	rootCmd.AddCommand(credentialCmd)
}
