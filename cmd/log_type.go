package cmd

import (
	"github.com/spf13/cobra"
)

var logsTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Manage BindPlane logs source types",
}

func init() {
	logsSourceCmd.AddCommand(logsTypeCmd)
}
