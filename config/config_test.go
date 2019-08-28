package config

import (
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

/*
TestListAccountsMissingFile: tests for missing/non-existing config file
-------------------------------------------
Produces an error message explaining that the file being read does
not exist
*/
func TestListAccountsMissingFile(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}

	err := ListAccounts()
	if err == nil {
		t.Errorf("An error should have occurred when reading a file that does not exist")
	}
}

/*
TestListAccountsFileEmpty: tests for any errors when a file is empty
Produces an error explaining that file being read does exist, but is empty.
*/
func TestListAccountsFileEmpty(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if err := ListAccounts(); err == nil {
		t.Errorf("An error should have occurred explaining that the file is empty")
	}
}

/*
TestListAccountsFileNotEmpty: tests for reading a file that is not empty
Displays a list of accounts that exist within the file that is being read.
*/
func TestListAccountsFileNotEmpty(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if err := ListAccounts(); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestAddAccountNoFileExists: tests for file creation if one is missing
This should create a new configuration file and add the account to the file
*/
func TestAddAccountNoFileExists(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}

	err := AddAccount("test4", "74a1aee1-ee9b-462f-88dc-65773eada2d7")
	if err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestAddAccountEmptyString: tests for empty name string
-------------------------------------------
Produces an error explaining that the name for an account must not be
empty, or just all spaces, it must have characters.
*/
func TestAddAccountEmptyString(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	err := AddAccount(" ", "74a1aee1-ee9b-462f-88dc-65773eada2d7")
	if err == nil {
		t.Errorf("An error should have been produced, saying the name string cannot be empty")
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestAddAccountFalseUUID: tests trying to add an account where UUID is not valid
Produces an error message explaining that the API Key given is not a valid UUID
*/
func TestAddAccountFalseUUID(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	err := AddAccount("test4", "74a1aee1-ee9b37489375%^^&%^&%")
	if err == nil {
		t.Errorf("Should produce an error message explaining the API Key is not" +
			" a valid UUID")
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestAddAccountDuplicateUUID: tests adding an account with an already
existing API Key
-------------------------------------------
Produces an error message saying that the API Key given already exists in the
configuration file
*/
func TestAddAccountDuplicateUUID(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	err := AddAccount("test4", "43c009e8-40f3-4506-93ef-411299cf4181")
	if err == nil {
		t.Errorf("Should produce an error explaining the API Key given already" +
			" exists in the configuration file")
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestAddAccountDuplicateName: tests adding an account with a name that
already exists
-------------------------------------------
Produces an error message saying the name given already exists in the
configuration file
*/
func TestAddAccountDuplicateName(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	err := AddAccount("test1", "74a1aee1-ee9b-462f-88dc-65773eada2d7")
	if err == nil {
		t.Errorf("This test should fail, duplicate names are not allowed.")
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestRemoveFileEmpty: tests the removing an account from an empty file
-------------------------------------------
Produces an error explaining that the account given does not match any accounts
that exist within the configuration file
*/
func TestRemoveFileEmpty(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if err := Remove("no-exist"); err == nil {
		t.Errorf("This file does exist, but the function should return an error")
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestRemoveExistingAccount: tests removing an existing account from the
configuration file defined by the user.
-------------------------------------------
Remove the account named 'test1' from the configuration file.
Return a nil error
*/
func TestRemoveExistingAccount(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if err := Remove("test1"); err != nil {
		t.Errorf(err.Error())
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestRemoveNonExistingAccount: tries to remove a non-existing account
-------------------------------------------
Produces an error message, explaining that the name is
not present in the list of accounts.
*/
func TestRemoveNonExistingAccount(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if err := Remove("no-exist"); err != nil {
		t.Errorf(err.Error())
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestSetCurrentExistingAccount: tests setting an account to current
-------------------------------------------
The account named 'test1' will be set as current
in the list of accounts.
*/
func TestSetCurrentExistingAccount(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if err := SetCurrent("test1"); err != nil {
		t.Errorf(err.Error())
	}
	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestSetCurrentNonExistingAccount: tests setting an account to current that does
not exist in the configuration file
-------------------------------------------
Produces an error message, explaining that the name is
not present in the configuration file.
*/
func TestSetCurrentNonExistingAccount(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if err := SetCurrent("no-exist"); err == nil {
		t.Errorf("Expected an error when trying to set the current account to an account that does not exist")
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestCurrentAPIKey
Case 1:
No environment variable exists.
No configuration file is present.
-------------------------------------------
Expected Behavior: Produce an error message
*/
func TestCurrentAPIKeyNoFileNoEnv(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Unsetenv("BINDPLANE_API_KEY")

	if _, err := CurrentAPIKey(); err == nil {
		t.Errorf("This should produce an error message explaining that neither the" +
			" configuration file, or environment variable, are present.")
	}
}

/*
TestCurrentAPIKeyNoFile
Case 2:
Environment variable is present
No configuration file is present
-------------------------------------------
Expected Behavior: The API Key should be set using the environment
variable that is defined, no errors.
*/
func TestCurrentAPIKeyNoFile(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "43c009e8-40f3-4506-93ef-411299cf4181")

	if _, err := CurrentAPIKey(); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestCurrentAPIKeyNoFileEmptyEnv no file, empty env variable exists
Case 2a:
Environment Variable is present
No configuration file exists
Environment variable being passed is an empty string
-------------------------------------------
Expected Behavior: Produce an error saying that the API Key is not set
*/
func TestCurrentAPIKeyNoFileEmptyEnv(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "")

	if _, err := CurrentAPIKey(); err == nil {
		t.Errorf("This should have produced an error saying the API Key is not set")
	}
}

/*
TestCurrentAPIKeyNoFileBadEnv no file, bad env variable exists
Case 2b:
Environment Variable is present
No configuration file exists
Environment variable being passed is not a valid UUID
-------------------------------------------
Expected Behavior: Produce an error saying that the API Key is
not a valid UUID
*/
func TestCurrentAPIKeyNoFileBadEnv(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "1")

	if _, err := CurrentAPIKey(); err == nil {
		t.Errorf("Should produce an error explaining the API Key is not a valid UUID")
	}
}

/*
TestCurrentAPIKeyNoEnv no env variable, file exists
Case 3:
Environment variable is not present
Configuration file exists
Configuration file has valid data
-------------------------------------------
Expected Behavior: The API Key will be set using the configuration file that
has been defined by the user.
*/
func TestCurrentAPIKeyNoEnv(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Unsetenv("BINDPLANE_API_KEY")
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if _, err := CurrentAPIKey(); err != nil {
		t.Errorf(err.Error())
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestCurrentAPIKeyNoEnvNoCurrent no env variable, file exists
Case 3a:
Environment variable is not present
Configuration file exists
Configuration file has valid data
Configuration file does not have an account with Current set to true
-------------------------------------------
Expected Behavior: Produce an error explaining that the configuration file is
present, but does not have an account set to current
*/
func TestCurrentAPIKeyNoEnvNoCurrent(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Unsetenv("BINDPLANE_API_KEY")
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	// test3 is set to Current, so remove it for the test
	Remove("test3")

	if _, err := CurrentAPIKey(); err == nil {
		t.Errorf("An error should have occurred when trying to retrieve the API Key")
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

/*
TestCurrentAPIKeyWarning env variable and file exist
Case 4:
Environment variable is present
Environment variable is valid
Configuration File is present
Configuration File has valid data
-------------------------------------------
Expected Behavior: Print a warning to Stderr explaining that the environment
variable will ALWAYS take precedence over the configuration file. Set the API
Key using the environment variable and return a nil error
*/
func TestCurrentAPIKeyWarning(t *testing.T) {

	os.Setenv("BINDPLANE_CONFIG_FILE", configFile)
	os.Setenv("BINDPLANE_API_KEY", "43c009e8-40f3-4506-93ef-411299cf4181")
	if err := createConfigFile(); err != nil {
		t.Errorf(err.Error())
	}
	if err := resetConfigFile(); err != nil {
		t.Errorf(err.Error())
	}

	if _, err := CurrentAPIKey(); err != nil {
		t.Errorf(err.Error())
	}

	if err := os.Remove(configFile); err != nil {
		t.Errorf(err.Error())
	}
}

// createConfigFile is a helper function that creates the test configuration file
func createConfigFile() error {
	emptyFile, err := os.Create(configFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	emptyFile.Close()
	return nil
}

// resetConfigFile adds content to the file using testFile
func resetConfigFile() error {
	testListBytes, err := json.Marshal(testFile)
	if err != nil {
		return err
	}

	return write(testListBytes)
}
