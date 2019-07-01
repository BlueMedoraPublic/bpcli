package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

var getCollectorGroupCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a collector group",
	Run: func(cmd *cobra.Command, args []string) {
		getCollectorGroup()
	},
}

func init() {
	collectorGroupCmd.AddCommand(getCollectorGroupCmd)
	getCollectorGroupCmd.Flags().StringVar(&groupID, "id", "", "The ID of the collector group")
	getCollectorGroupCmd.MarkFlagRequired("id")
}

func getCollectorGroup() {
	c, err := bp.GetCollectorGroup(groupID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := c.Print(jsonFmt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
