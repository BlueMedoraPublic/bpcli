package cmd

import (
	"os"
	"fmt"

	"bpcli/bindplane"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bindplane CLI version:", bindplane.VERSION)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
