package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var bashCompletion = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion",
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(bashCompletion)
}
