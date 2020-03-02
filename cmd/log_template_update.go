package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BlueMedoraPublic/bpcli/bindplane/sdk"

	"github.com/spf13/cobra"
)

var logTemplateUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a log template",
	Run: func(cmd *cobra.Command, args []string) {
		if err := updateLogTemplate(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logTemplateCmd.AddCommand(logTemplateUpdateCmd)
}

func updateLogTemplate() error {
	f, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	t := sdk.LogTemplate{}
	if err := json.Unmarshal(f, &t); err != nil {
		return err
	}

	if err := bp.UpdateLogTemplate(logTemplateID, t); err != nil {
		return err
	}

	fmt.Println("template updated")
	return nil
}
