package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var logTemplateDelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a log template",
	Run: func(cmd *cobra.Command, args []string) {
		if err := delLogTemplate(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	logTemplateCmd.AddCommand(logTemplateDelCmd)
}

func delLogTemplate() error {
	err := bp.DeleteLogTemplate(logTemplateID)
	if err != nil {
		return err
	}

	fmt.Println("template " + logTemplateID + " deleted")
	return nil
}
