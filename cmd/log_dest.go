package cmd

import (
	"github.com/spf13/cobra"
)

var logsDestCmd = &cobra.Command{
	Use:   "destination",
	Short: "Manage BindPlane log destinations",
}

func init() {
	logsCmd.AddCommand(logsDestCmd)
}
