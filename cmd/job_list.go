package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listJobCmd = &cobra.Command{
	Use:   "list",
	Short: "List all jobs",
	Run: func(cmd *cobra.Command, args []string) {
		listJobs()
	},
}

func init() {
	jobCmd.AddCommand(listJobCmd)
}

func listJobs() {
	j, err := bp.ListJobs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	for _, job := range j {
		if err := job.Print(false); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}
}
