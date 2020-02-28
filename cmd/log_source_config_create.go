package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var logSourceConfigCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a source config",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createLogSourceConfigs(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logsConfigCmd.AddCommand(logSourceConfigCreateCmd)
}

func createLogSourceConfigs() error {
	f, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}


	c, err := bp.CreateLogSourceConfig(f)
	if err != nil {
		return err
	}

	return c.Print(jsonFmt)
}
