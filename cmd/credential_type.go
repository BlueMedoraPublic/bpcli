package cmd

import (
	"github.com/spf13/cobra"
)

var credentialTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Manage credential type",
}

func init() {
	credentialCmd.AddCommand(credentialTypeCmd)
}
