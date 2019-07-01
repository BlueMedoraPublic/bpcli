package cmd

import (
	"os"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var createCredentialCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a credential",
	Run: func(cmd *cobra.Command, args []string) {
		createCredential()
	},
}

func init() {
	credentialCmd.AddCommand(createCredentialCmd)
	createCredentialCmd.Flags().StringVar(&credentialFile, "file", "", "The credential json file")
	createCredentialCmd.MarkFlagRequired("file")
}

func createCredential() {
	// read the json file
	f, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	// post the payload to api, creating the source
	resp, err := bp.CreateCredential(f)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := resp.Print(jsonFmt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
