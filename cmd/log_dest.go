package cmd

import (
	"github.com/spf13/cobra"
)

var logsDestCmd = &cobra.Command{
	Use:   "destination",
	Short: "Manage BindPlane log destinations",
}

var logsDestTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Manage BindPlane log destination types",
}

var logsDestConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage BindPlane log destination configurations",
}

func init() {
	logsCmd.AddCommand(logsDestCmd)
	logsDestCmd.AddCommand(logsDestTypeCmd)
	logsDestCmd.AddCommand(logsDestConfigCmd)

	// define all destination sub command flags here
	const (
		logDestIDFlagName    = "config-id"
		logDestIDFlagShort   = ""
		logDestIDFlagDefault = ""
		logDestIDFlagDesc    = "destination id"

		logDestTypeIDFlagName    = "type-id"
		logDestTypeIDFlagShort   = ""
		logDestTypeIDFlagDefault = ""
		logDestTypeIDFlagDesc    = "destination type id"
	)

	logDestConfigCreateCmd.Flags().StringVarP(&sourceFile, "file", "", "", "destination json file")
	logDestConfigCreateCmd.MarkFlagRequired("file")

	logDestConfigDeleteCmd.Flags().StringVarP(&logConfigID, logDestIDFlagName, logDestIDFlagShort, logDestIDFlagDefault, logDestIDFlagDesc)
	logDestConfigDeleteCmd.MarkFlagRequired(logDestIDFlagName)

	logDestConfigGetCmd.Flags().StringVarP(&logConfigID, logDestIDFlagName, logDestIDFlagShort, logDestIDFlagDefault, logDestIDFlagDesc)
	logDestConfigGetCmd.MarkFlagRequired(logDestIDFlagName)

	logDestConfigUpdateVersion.Flags().StringVarP(&logConfigID, logDestIDFlagName, logDestIDFlagShort, logDestIDFlagDefault, logDestIDFlagDesc)
	logDestConfigUpdateVersion.MarkFlagRequired(logDestIDFlagName)

	logDestTypeGetCmd.Flags().StringVarP(&logTypeID, logDestTypeIDFlagName, logDestTypeIDFlagShort, logDestTypeIDFlagDefault, logDestTypeIDFlagDesc)
	logDestTypeGetCmd.MarkFlagRequired(logDestTypeIDFlagName)

	logDestTypeParametersCmd.Flags().StringVarP(&logTypeID, logDestTypeIDFlagName, logDestTypeIDFlagShort, logDestTypeIDFlagDefault, logDestTypeIDFlagDesc)
	logDestTypeParametersCmd.MarkFlagRequired(logDestTypeIDFlagName)
}
