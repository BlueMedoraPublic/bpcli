package cmd

import (
	"github.com/spf13/cobra"
)

var logsSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage BindPlane log sources",
}

var logsTypeCmd = &cobra.Command{
	Use:   "type",
	Short: "Manage BindPlane logs source types",
}

var logsConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage BindPlane logs source configurations",
}

func init() {
	logsCmd.AddCommand(logsSourceCmd)
	logsSourceCmd.AddCommand(logsTypeCmd)
	logsSourceCmd.AddCommand(logsConfigCmd)

	// define all agent source sub command flags here
	const (
		logSourceIDFlagName    = "source-id"
		logSourceIDFlagShort   = ""
		logSourceIDFlagDefault = ""
		logSourceIDFlagDesc    = ""

		logSourceTypeIDFlagName    = "source-type-id"
		logSourceTypeIDFlagShort   = ""
		logSourceTypeIDFlagDefault = ""
		logSourceTypeIDFlagDesc    = "source type id"
	)

	logSourceConfigCreateCmd.Flags().StringVarP(&sourceFile, "file", "", "", "The source json file")
	logSourceConfigCreateCmd.MarkFlagRequired("file")

	logSourceConfigDeleteCmd.Flags().StringVarP(&logConfigID, logSourceIDFlagName, logSourceIDFlagShort, logSourceIDFlagDefault, logSourceIDFlagDesc)
	logSourceConfigDeleteCmd.MarkFlagRequired(logSourceIDFlagName)

	logSourceConfigGetCmd.Flags().StringVarP(&logConfigID, logSourceIDFlagName, logSourceIDFlagShort, logSourceIDFlagDefault, logSourceIDFlagDesc)
	logSourceConfigGetCmd.MarkFlagRequired(logSourceIDFlagName)

	logSourceConfigUpdateVersion.Flags().StringVarP(&logConfigID, logSourceIDFlagName, logSourceIDFlagShort, logSourceIDFlagDefault, logSourceIDFlagDesc)
	logSourceConfigUpdateVersion.MarkFlagRequired(logSourceIDFlagName)

	logSourceTypeGetCmd.Flags().StringVarP(&logConfigID, logSourceTypeIDFlagName, logSourceTypeIDFlagShort, logSourceTypeIDFlagDefault, logSourceTypeIDFlagDesc)
	logSourceTypeGetCmd.MarkFlagRequired(logSourceTypeIDFlagName)

	logSourceTypeParametersCmd.Flags().StringVarP(&logConfigID, logSourceTypeIDFlagName, logSourceTypeIDFlagShort, logSourceTypeIDFlagDefault, logSourceTypeIDFlagDesc)
	logSourceTypeParametersCmd.MarkFlagRequired(logSourceTypeIDFlagName)
}
