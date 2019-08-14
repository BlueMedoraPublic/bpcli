package config

import (
	//"os/exec"
	"log"
	"os"
	"testing"
)

const configFile = "/tmp/.bpcli"

var testFile = []account{
	{
		Name:    "test1",
		Key:     "43c009e8-40f3-4506-93ef-411299cf4181",
		Current: false,
	},
	{
		Name:    "test2",
		Key:     "024f5f47-e0e1-4f61-bf9c-39976db7f4a8",
		Current: false,
	},
	{
		Name:    "test3",
		Key:     "91f5e858-eb9f-4136-8a37-bedeebbb7885",
		Current: true,
	},
}

//Writes tests for all functions within config.go

//Will need to execute a bash script to set the file being tested back to
//it's original state
//var delScript = exec.Command("/bin/sh", dir+"/test-file-delete")
//var resetScript = exec.Command("/bin/sh", dir+"/test-file-reset")

// TestListAccounts tests for missing/non-existing config file
func TestListAccounts(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()

	err := ListAccounts()
	if err == nil {
		t.Errorf(err.Error())
	}
}

// TestListAccounts2 tests for any errors when a file exists
/*func TestListAccounts2(t *testing.T) {

	err := ListAccounts()
	if err != nil {
		t.Errorf("The function should not be returning an error if file exists")
	}
}*/

// TestRemove tests the remove function in config.go
// func TestRemove(t *testing.T) {
// 	err := Remove("no-exist")
// 	if err != nil {
// 		t.Errorf("This file does not exist, but the function should not return an error")
// 	}
// }

func createConfigFile() {
	emptyFile, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	emptyFile.Close()
}

func cleanconfigFile() {

}
