package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logTemplateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List log templates",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listLogTemplate(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logTemplateCmd.AddCommand(logTemplateListCmd)
}

func listLogTemplate() error {
	t, err := bp.ListLogTemplates()
	if err != nil {
		return err
	}

	for _, t := range t {
		if err := t.Print(jsonFmt); err != nil {
			return err
		}
	}
	return nil
}
