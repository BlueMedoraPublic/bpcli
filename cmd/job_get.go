package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var getJobCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a job",
	Run: func(cmd *cobra.Command, args []string) {
		getJob()
	},
}

func init() {
	jobCmd.AddCommand(getJobCmd)
	getJobCmd.Flags().StringVar(&jobID, "id", "", "The ID of the job")
	getJobCmd.MarkFlagRequired("id")
}

func getJob() {
	j, err := bp.GetJob(jobID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := j.Print(jsonFmt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
