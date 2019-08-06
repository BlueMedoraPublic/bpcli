package cmd

import (
	"github.com/spf13/cobra"
)

var listAccountCmd = &cobra.Command{
	Use:   "list",
	Short: "List available BindPlane accounts",
}

func init() {
	accountCmd.AddCommand(listAccountCmd)
}
