package cmd

import (
	"fmt"
	"os"

	"github.com/BlueMedoraPublic/bpcli/config"
	"github.com/spf13/cobra"
)

var listAccountCmd = &cobra.Command{
	Use:   "list",
	Short: "List available BindPlane accounts",
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	accountCmd.AddCommand(listAccountCmd)
}

func list() {
	if err := config.ListAccounts(); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
