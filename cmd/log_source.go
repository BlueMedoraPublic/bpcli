package cmd

import (
	"github.com/spf13/cobra"
)

var logsSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage BindPlane log sources",
}

func init() {
	logsCmd.AddCommand(logsSourceCmd)
}
