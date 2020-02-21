package cmd

import (
	"github.com/spf13/cobra"
)

var logsConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage BindPlane logs source configurations",
}

func init() {
	logsSourceCmd.AddCommand(logsConfigCmd)
}
