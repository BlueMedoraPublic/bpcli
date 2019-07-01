package cmd

import (
	"github.com/spf13/cobra"
)

var collectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "Manage collectors and collector groups",
}

func init() {
	rootCmd.AddCommand(collectorCmd)
}

// setGroupID will set groupID equal to collectorID if
// groupID is not set with a command line argument
func setGroupID() {
	if len(groupID) == 0 {
		groupID = collectorID
	}
}
