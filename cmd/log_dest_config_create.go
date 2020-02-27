package cmd

import (
	"fmt"
	"os"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var logDestConfigCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a destination config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createLogDestConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsDestConfigCmd.AddCommand(logDestConfigCreateCmd)
	logDestConfigCreateCmd.Flags().StringVar(&sourceFile, "file", "", "The destination json file")
	logDestConfigCreateCmd.MarkFlagRequired("file")
}

func createLogDestConfigs() error {
	f, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	d, err := bp.CreateLogDestConfig(f)
	if err != nil {
		return err
	}

	return d.Print(jsonFmt)
}
