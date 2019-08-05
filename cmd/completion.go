package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion",
	Run: func(cmd *cobra.Command, args []string) {
		genTabCompletion()
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.Flags().BoolVar(&zshCompletion, "zsh", false, "generates zsh tab completion script")
}

func genTabCompletion() {
	if zshCompletion {
		genZshCompletion()
	} else {
		genBashCompletion()
	}
}

func genBashCompletion() {
	rootCmd.GenBashCompletion(os.Stdout)
}

func genZshCompletion() {
	rootCmd.GenZshCompletion(os.Stdout)
}
