package cmd

import (
	"fmt"
	"os"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var logSourceConfigCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Describe a source config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createLogSourceConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsConfigCmd.AddCommand(logSourceConfigCreateCmd)
	logSourceConfigCreateCmd.Flags().StringVar(&sourceFile, "file", "", "The source json file")
	logSourceConfigCreateCmd.MarkFlagRequired("file")
}

func createLogSourceConfigs() error {
	f, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	var b []byte
	b, err = bp.CreateLogSourceConfig(f)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
