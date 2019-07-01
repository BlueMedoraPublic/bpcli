package cmd

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"

	"github.com/BlueMedoraPublic/bpcli/bindplane/sdk"

	"github.com/spf13/cobra"
)

// createSourceCmd represents the create command
var createSourceCmd = &cobra.Command{
	Use:   "create",
	Short: "Configure a new source",
	Run: func(cmd *cobra.Command, args []string) {
		createSource()
	},
}

func init() {
	sourceCmd.AddCommand(createSourceCmd)
	createSourceCmd.Flags().StringVar(&sourceFile, "file", "", "The source json file")
	createSourceCmd.MarkFlagRequired("file")
}

func createSource() {
	// read the json file
	f, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	/*
	NOTE: we are sending 'f' ([]byte) to the API, only convert
	to a struct to perform validation on the object
	*/
	var s sdk.SourceConfigCreate
    if err := json.Unmarshal(f, &s); err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
    }
	if err := s.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	r, err := bp.CreateSource(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := r.Print(jsonFmt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
