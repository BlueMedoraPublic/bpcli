package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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

	t := sdk.LogTemplate{}
	if err := json.Unmarshal(f, &t); err != nil {
		return err
	}

	template, err := bp.CreateLogTemplate(t)
	if err != nil {
		return err
	}

	fmt.Println("template created: " + template.ID)
	return nil
}
