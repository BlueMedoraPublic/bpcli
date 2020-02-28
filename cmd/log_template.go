package cmd

import (
	"github.com/spf13/cobra"
)

var logTemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage BindPlane log templates",
}

func init() {
	logsCmd.AddCommand(logTemplateCmd)

	// define all agent sub command flags here
	const (
		logTemplateIDFlagName    = "template-id"
		logTemplateIDFlagShort   = ""
		logTemplateIDFlagDefault = ""
		logTemplateIDFlagDesc    = "log template id"
	)

	logTemplateGetCmd.Flags().StringVarP(&logTemplateID, logTemplateIDFlagName, logTemplateIDFlagShort, logTemplateIDFlagDefault, logTemplateIDFlagDesc)
	logTemplateGetCmd.MarkFlagRequired(logTemplateIDFlagName)

	logTemplateDelCmd.Flags().StringVarP(&logTemplateID, logTemplateIDFlagName, logTemplateIDFlagShort, logTemplateIDFlagDefault, logTemplateIDFlagDesc)
	logTemplateDelCmd.MarkFlagRequired(logTemplateIDFlagName)

	logTemplateCreateCmd.Flags().StringVarP(&sourceFile, "file", "", "", "template json file")
	logTemplateCreateCmd.MarkFlagRequired("file")

	logTemplateUpdateCmd.Flags().StringVarP(&sourceFile, "file", "", "", "template json file")
	logTemplateUpdateCmd.Flags().StringVarP(&logTemplateID, logTemplateIDFlagName, logTemplateIDFlagShort, logTemplateIDFlagDefault, logTemplateIDFlagDesc)
	logTemplateUpdateCmd.MarkFlagRequired(logTemplateIDFlagName)
	logTemplateUpdateCmd.MarkFlagRequired("file")
}
