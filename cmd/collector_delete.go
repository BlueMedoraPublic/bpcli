package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var delCollectorCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a collector",
	Run: func(cmd *cobra.Command, args []string) {
		delCollector()
	},
}

func init() {
	collectorCmd.AddCommand(delCollectorCmd)
	delCollectorCmd.Flags().StringVar(&groupID, "group-id", "", "Group ID (default to collector id if not passed)" )
	delCollectorCmd.Flags().StringVar(&collectorID, "id", "", "Collector ID")
	delCollectorCmd.MarkFlagRequired("id")
}

func delCollector() {
	setGroupID()

	err := bp.DeleteCollector(groupID + "/" + collectorID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("collector deleted:", groupID + "/" + collectorID)
}
