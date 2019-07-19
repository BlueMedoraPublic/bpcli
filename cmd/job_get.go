package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BlueMedoraPublic/bpcli/bindplane/sdk"

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
	getJobCmd.Flags().BoolVarP(&watch, "watch", "", false, "Monitor the job's status in the forground until it has passed or failed.")
	getJobCmd.MarkFlagRequired("id")
}

func getJob() {
	for {
		j, err := bp.GetJob(jobID)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		if err := j.Print(jsonFmt); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		// exit if watch not passed
		if watch == false {
			os.Exit(0)
		}

		// exit if job is finished
		if jobFinished(j) == true {
			os.Exit(0)
		}
		time.Sleep(time.Second * 5)
	}
}

// return true if job is complete (passed or failed)
func jobFinished(j sdk.Job) bool {
	if strings.ToLower(j.Status) == "in progress" {
		return false
	}
	return true
}
