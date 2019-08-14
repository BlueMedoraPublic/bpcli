package config

import (
	//"os/exec"
	"encoding/json"
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

// TestListAccounts tests for missing/non-existing config file
func TestListAccounts(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Remove(configFile)

	err := ListAccounts()
	if err == nil {
		t.Errorf(err.Error())
	}
}

// TestListAccounts2 tests for any errors when a file is empty
func TestListAccounts2(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()

	err := ListAccounts()
	if err == nil {
		t.Errorf(err.Error())
	}
}

// TestListAccounts3 tests for reading a file that is not empty
func TestListAccounts3(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := ListAccounts()
	if err != nil {
		t.Errorf(err.Error())
	}
}

// TestAddAccount tests for file creation if one is missing
func TestAddAccount(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Remove(configFile)

	err := AddAccount("test4", "74a1aee1-ee9b-462f-88dc-65773eada2d7")
	if err != nil {
		t.Errorf(err.Error())
	}
}

// TestAddAccount2 tests for empty name string
func TestAddAccount2(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := AddAccount("", "74a1aee1-ee9b-462f-88dc-65773eada2d7")
	if err == nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestAddAccount3
func TestAddAccount3(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := AddAccount("test4", "74a1aee1-ee9b37489375%^^&%^&%")
	if err == nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestAddAccount4
func TestAddAccount4(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := AddAccount("test4", "43c009e8-40f3-4506-93ef-411299cf4181")
	if err == nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestAddAccount5
func TestAddAccount5(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := AddAccount("test1", "74a1aee1-ee9b-462f-88dc-65773eada2d7")
	if err == nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestRemove tests the removing an account from empty file
func TestRemove(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()

	err := Remove("no-exist")
	if err == nil {
		t.Errorf("This file does exist, but the function should return an error")
	}

	os.Remove(configFile)
}

// TestRemove2
func TestRemove2(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := Remove("test1")
	if err != nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestRemove3
func TestRemove3(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := Remove("no-exist")
	if err != nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestSetCurrent sets an account to current that exists
func TestSetCurrent(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := SetCurrent("test1")
	if err != nil {
		t.Errorf(err.Error())
	}
	os.Remove(configFile)
}

// TestSetCurrent2
func TestSetCurrent2(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	createConfigFile()
	resetConfigFile()

	err := SetCurrent("no-exist")
	if err != nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestCurrentAPIKey no file, no env variable
func TestCurrentAPIKey(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Unsetenv("BINDPLANE_API_KEY")

	_, err := CurrentAPIKey()
	if err == nil {
		t.Errorf(err.Error())
	}
}

// TestCurrentAPIKey2 no file, env variable exists
func TestCurrentAPIKey2(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "43c009e8-40f3-4506-93ef-411299cf4181")

	_, err := CurrentAPIKey()
	if err != nil {
		t.Errorf(err.Error())
	}
}

// TestCurrentAPIKey3 no file, bad env variable exists
func TestCurrentAPIKey3(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "")

	_, err := CurrentAPIKey()
	if err == nil {
		t.Errorf(err.Error())
	}
}

// TestCurrentAPIKey4 no file, bad env variable exists
func TestCurrentAPIKey4(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "1")

	_, err := CurrentAPIKey()
	if err == nil {
		t.Errorf(err.Error())
	}
}

// TestCurrentAPIKey5 no env variable, file exists
func TestCurrentAPIKey5(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Unsetenv("BINDPLANE_API_KEY")
	createConfigFile()
	resetConfigFile()

	_, err := CurrentAPIKey()
	if err != nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestCurrentAPIKey6 no env variable, file exists, no current
func TestCurrentAPIKey6(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Unsetenv("BINDPLANE_API_KEY")
	createConfigFile()
	resetConfigFile()

	Remove("test3")

	_, err := CurrentAPIKey()
	if err == nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

// TestCurrentAPIKey7 env variable and file exist
func TestCurrentAPIKey7(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "43c009e8-40f3-4506-93ef-411299cf4181")
	createConfigFile()
	resetConfigFile()

	_, err := CurrentAPIKey()
	if err != nil {
		t.Errorf(err.Error())
	}

	os.Remove(configFile)
}

func createConfigFile() {
	emptyFile, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	emptyFile.Close()
}

func resetConfigFile() error {

	testListBytes, err := json.Marshal(testFile)
	if err != nil {
		return err
	}

	return write(testListBytes)
}
