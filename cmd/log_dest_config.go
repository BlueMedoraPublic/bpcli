package cmd

import (
	"github.com/spf13/cobra"
)

var logsDestConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage BindPlane log destination configurations",
}

func init() {
	logsDestCmd.AddCommand(logsDestConfigCmd)
}
