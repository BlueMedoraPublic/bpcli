package cmd

import (
	"github.com/spf13/cobra"
)

var logsDestTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Manage BindPlane log destination types",
}

func init() {
	logsDestCmd.AddCommand(logsDestTypeCmd)
}
