package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

	"github.com/BlueMedoraPublic/bpcli/bindplane/sdk"

	"github.com/spf13/cobra"
)

var logTemplateCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a log template",
	Run: func(cmd *cobra.Command, args []string) {
		if err := createLogTemplate(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logTemplateCmd.AddCommand(logTemplateCreateCmd)
}

func createLogTemplate() error {
	f, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	t := sdk.LogTemplateCreate{}
	if err := json.Unmarshal(f, &t); err != nil {
		return err
	}

	if err := bp.CreateLogTemplate(t); err != nil {
		return err
	}

	fmt.Println("template created")
	return nil
}
