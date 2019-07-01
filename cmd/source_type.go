package cmd

import (
	"github.com/spf13/cobra"
)

var sourceTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Manage source types",
}

func init() {
	sourceCmd.AddCommand(sourceTypeCmd)
}
