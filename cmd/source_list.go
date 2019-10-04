package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listSourceCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured sources",
	Run: func(cmd *cobra.Command, args []string) {
		if err := listSources(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	sourceCmd.AddCommand(listSourceCmd)
}

func listSources() error {
	s, err := bp.GetSources()
	if err != nil {
		return err
	}

	if jsonFmt == false {
		for _, source := range s {
			if err := source.Print(false); err != nil {
				return err
			}
		}
		return nil
	}

	// instead of printing each source object on its own,
	// marshal them into a json array to make it valid json
	x, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(x))
	return nil
}
