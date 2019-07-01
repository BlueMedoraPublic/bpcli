package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var deleteSourceCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a source",
	Run: func(cmd *cobra.Command, args []string) {
		deleteSource()
	},
}

func init() {
	sourceCmd.AddCommand(deleteSourceCmd)
	deleteSourceCmd.Flags().StringVar(&sourceID, "id", "", "The ID of the source")
	deleteSourceCmd.MarkFlagRequired("id")
}

func deleteSource() {
	_, err := bp.DeleteSource(sourceID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("deleted source", sourceID)
}
