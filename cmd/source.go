package cmd

import (
	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage sources",
}

func init() {
	rootCmd.AddCommand(sourceCmd)
}
