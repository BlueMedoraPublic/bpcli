package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logTemplateGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Desribe a log template",
	Run: func(cmd *cobra.Command, args []string) {
		if err := getLogTemplate(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logTemplateCmd.AddCommand(logTemplateGetCmd)
}

func getLogTemplate() error {
	t, err := bp.GetLogTemplate(logTemplateID)
	if err != nil {
		return err
	}

	return t.Print(jsonFmt)
}
