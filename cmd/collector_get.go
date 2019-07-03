package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var getCollectorCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a collector",
	Run: func(cmd *cobra.Command, args []string) {
		getCollector()
	},
}

func init() {
	collectorCmd.AddCommand(getCollectorCmd)
	getCollectorCmd.Flags().StringVar(&collectorID, "id", "", "The ID of the collector")
	getCollectorCmd.Flags().StringVar(&groupID, "group-id", "", "The Group ID the collector belongs to (Defaults to '--id')")
	getCollectorCmd.MarkFlagRequired("id")
}

func getCollector() {
	setGroupID()

	c, err := bp.GetCollector(groupID + "/" + collectorID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := c.Print(jsonFmt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
