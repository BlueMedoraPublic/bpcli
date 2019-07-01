package cmd

import (
	"github.com/spf13/cobra"
)

var collectorGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage collector groups",
}

func init() {
	collectorCmd.AddCommand(collectorGroupCmd)
}
