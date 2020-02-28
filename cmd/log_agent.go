package cmd

import (
	"github.com/spf13/cobra"
)

var logAgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Manage BindPlane log agents",
}

var logAgentDestinationCmd = &cobra.Command{
	Use:   "destination",
	Short: "Manage log destinations",
}

var logAgentSourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage log sources",
}

var logAgentTaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage log agent tasks",
}

func init() {
	logsCmd.AddCommand(logAgentCmd)
	logAgentCmd.AddCommand(logAgentDestinationCmd)
	logAgentCmd.AddCommand(logAgentSourceCmd)
	logAgentCmd.AddCommand(logAgentTaskCmd)

	// define all agent sub command flags here
	const (
		agentIDFlagName = "agent-id"
		agentIDFlagShort = "a"
		agentIDFlagDefault  = ""
		agentIDFlagDesc    = "log agent id"

		destIDFlagName  = "destination-id"
		destIDFlagShort  = ""
		destIDFlagDefault  = ""
		destIDFlagDesc  = "destination configuration ID"

		sourceIDFlagName = "source-id"
		sourceIDFlagShort = ""
		sourceIDFlagDefault = ""
		sourceIDFlagDesc = "source config id"

		taskIDFlagName = "task-id"
		taskIDFlagShort = ""
		taskIDFlagDefault = ""
		taskIDFlagDesc = "task id"
	)

	logAgentDeleteCmd.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentDeleteCmd.MarkFlagRequired(agentIDFlagName)

	logAgentDestDel.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentDestDel.Flags().StringVarP(&logConfigID, destIDFlagName, destIDFlagShort, destIDFlagDefault, destIDFlagDesc)
	logAgentDestDel.MarkFlagRequired(agentIDFlagName)
	logAgentDestDel.MarkFlagRequired(destIDFlagName)

	logAgentDestDeploy.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentDestDeploy.Flags().StringVarP(&logConfigID, destIDFlagName, destIDFlagShort, destIDFlagDefault, destIDFlagDesc)
	logAgentDestDeploy.MarkFlagRequired(agentIDFlagName)
	logAgentDestDeploy.MarkFlagRequired(destIDFlagName)

	logAgentDestGet.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentDestGet.Flags().StringVarP(&logConfigID, destIDFlagName, destIDFlagShort, destIDFlagDefault, destIDFlagDesc)
	logAgentDestGet.MarkFlagRequired(agentIDFlagName)
	logAgentDestGet.MarkFlagRequired(destIDFlagName)

	logAgentDestList.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentDestList.MarkFlagRequired(agentIDFlagName)

	logAgentGetCmd.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentGetCmd.MarkFlagRequired(agentIDFlagName)

	logAgentSourceDelete.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentSourceDelete.Flags().StringVarP(&logConfigID, sourceIDFlagName, sourceIDFlagShort, sourceIDFlagDefault, sourceIDFlagDesc)
	logAgentSourceDelete.MarkFlagRequired(agentIDFlagName)
	logAgentSourceDelete.MarkFlagRequired(sourceIDFlagName)

	logAgentSourceDeploy.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentSourceDeploy.Flags().StringVarP(&logConfigID, sourceIDFlagName, sourceIDFlagShort, sourceIDFlagDefault, sourceIDFlagDesc)
	logAgentSourceDeploy.MarkFlagRequired(agentIDFlagName)
	logAgentSourceDeploy.MarkFlagRequired(sourceIDFlagName)

	logAgentSourceGet.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentSourceGet.Flags().StringVarP(&logConfigID, sourceIDFlagName, sourceIDFlagShort, sourceIDFlagDefault, sourceIDFlagDesc)
	logAgentSourceGet.MarkFlagRequired(agentIDFlagName)
	logAgentSourceGet.MarkFlagRequired(sourceIDFlagName)

	logAgentTaskGetCmd.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentTaskGetCmd.Flags().StringVarP(&logTaskID, taskIDFlagName, taskIDFlagShort, taskIDFlagDefault, taskIDFlagDesc)
	logAgentTaskGetCmd.MarkFlagRequired(agentIDFlagName)
	logAgentTaskGetCmd.MarkFlagRequired(taskIDFlagName)

	logAgentSourceList.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentSourceList.MarkFlagRequired(agentIDFlagName)

	logAgentUpdateVersion.Flags().StringVarP(&logAgentID, agentIDFlagName, agentIDFlagShort, agentIDFlagDefault, agentIDFlagDesc)
	logAgentUpdateVersion.MarkFlagRequired(agentIDFlagName)
}
