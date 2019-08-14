package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/BlueMedoraPublic/bpcli/util/uuid"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

// account stores users BindPlane account information
type account struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	Current bool   `json:"current"`
}

// ListAccounts prints a formatted list of users read from the configuration file
func ListAccounts() error {

	currentList, err := read()
	if err != nil {
		return err
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	if len(currentList) == 0 {
		return errors.New(path + " is empty! try adding a new account to the list\n")
	}

	fmt.Println("List of Account Names. * Denotes Current Account")

	// Print the list in a formatted way
	for _, acc := range currentList {
		if acc.Current == true {
			fmt.Println("* " + acc.Name)
		} else {
			fmt.Println(acc.Name)
		}
	}
	return nil
}

// AddAccount appends an account to the configuration file
func AddAccount(name string, key string) error {

	currentList, err := read()
	if err != nil {
		os.Stderr.WriteString("ERROR: " + err.Error() + "\n")

		path, err := configPath()
		if err != nil {
			return err
		}

		emptyFile, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		os.Stderr.WriteString("Creating a new file at: " + path + "\n")
		emptyFile.Close()
	}

	currentList, err = read()
	if err != nil {
		return err
	}

	if len(name) == 0 {
		return errors.New("The name cannot be an empty string")
	}

	if !uuid.IsUUID(key) {
		return errors.New("The API Key given is not a valid UUID")
	}

	b, err := uniqueUUID(key)
	if err != nil {
		return err
	}
	if b == false {
		return errors.New("The API Key given already exists within the config file")
	}

	n, err := uniqueName(name)
	if err != nil {
		return err
	}
	if n == false {
		return errors.New("The name given already exists within the config file")
	}

	a := account{Name: name, Key: key, Current: false}
	newList := append(currentList, a)

	newListBytes, err := json.Marshal(newList)
	if err != nil {
		return err
	}

	return write(newListBytes)
}

// Remove erases an account from the configuration file
func Remove(name string) error {

	currentList, err := read()
	if err != nil {
		return err
	}

	newList := currentList

	if !(len(newList) > 0) {
		return errors.New("The account list exists, but it is empty")
	}

	for i := 0; i < (len(newList)); i++ {
		if name != newList[i].Name {
			continue
		} else {
			newList = append(newList[:i], newList[i+1:]...)
			break
		}
	}

	if cmp.Equal(newList, currentList) {
		fmt.Println("No names matched the given input\n" +
			"Name Given: " + name + "\n")
		return nil
	}

	newListBytes, err := json.Marshal(newList)
	if err != nil {
		return err
	}

	return write(newListBytes)
}

// GetCurrentAPIKey attempts to retrieve an API key
func GetCurrentAPIKey() (string, error) {

	errMsg := `To use bpcli you must do either of the following:
	1. Set an environment variable named BINDPLANE_API_KEY
	2. Define a configuration file`

	// Check environment variable
	apiKey, envExists := os.LookupEnv("BINDPLANE_API_KEY")

	fileExists, _ := checkConfig()

	// No env variable, no file
	// Error
	if !fileExists && !envExists {
		return apiKey, errors.New("ERROR: The BINDPLANE_API_KEY environment variable is not present\n" +
			"ERROR: The configuration file is not present\n" +
			errMsg)
	}

	// Env variable exists, no file
	if !fileExists && envExists {
		if len(apiKey) <= 0 {
			return apiKey, errors.New("ERROR: The environment variable is not set\n" +
				errMsg)
		}
		if !uuid.IsUUID(apiKey) {
			return apiKey, errors.New("ERROR: The API Key given is not a valid UUID\n" +
				errMsg)
		}
		return apiKey, nil
	}

	// No env variable, config file w/o current set
	// Error
	if !envExists && fileExists {
		b, err := hasCurrent()
		if err != nil {
			return apiKey, err
		}
		if b == false {
			return apiKey, errors.New("ERROR: An environment variable is not present.\n" +
				"ERROR: A configuration file exists, but does not have an account set to current.\n" +
				errMsg + "\n" +
				"use the `bpcli account set` command to set an account to current in the config file.")
		}
	}

	// No env variable, config file w/ current set
	if !envExists && fileExists {
		b, err := hasCurrent()
		if err != nil {
			return apiKey, err
		}
		if b == true {
			return getCurrentFromConfig()
		}
	}

	// Env variables exists, file exists
	if envExists && fileExists {
		os.Stderr.WriteString("WARNING: An environment variable is set and a configuration file exists\n" +
			"The environment variable will ALWAYS take precedence over the configuration file\n" +
			"If you would like to use the configuration file, remove the environment variable\n" +
			"****COLLECTOR LIST****\n")

		return apiKey, nil
	}

	return apiKey, nil
}

// SetCurrent sets a chosen account to be the current account being worked in
func SetCurrent(name string) error {

	currentList, err := read()
	if err != nil {
		return err
	}

	b, err := accountExists(name)
	if err != nil {
		return err
	}
	if b == true {
		for i := range currentList {
			if name == currentList[i].Name {
				currentList[i].Current = true
			} else {
				currentList[i].Current = false
			}
		}

		updatedListBytes, err := json.Marshal(currentList)
		if err != nil {
			return err
		}

		return write(updatedListBytes)
	}

	return nil
}

// read Returns an array of accounts read from the configuration file
func read() ([]account, error) {
	accountList := []account{}

	filePath, err := configPath()
	if err != nil {
		return accountList, err
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return accountList, err
	}
	if len(file) == 0 {
		return accountList, nil
	}

	return accountList, json.Unmarshal(file, &accountList)
}

// write is a helper function that will write/re-write the configuration file
func write(list []byte) error {

	filePath, err := configPath()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, list, 0644)
}

// uniqueUUID checks the account list for duplicate UUIDs
func uniqueUUID(key string) (bool, error) {

	currentList, err := read()
	if err != nil {
		return false, err
	}

	if uuid.IsUUID(key) {
		for _, acc := range currentList {
			if key == acc.Key {
				return false, nil
			}
		}
	} else {
		return false, errors.New("The value given was not a valid UUID")
	}

	return true, nil
}

// uniqueName checks the users account list for duplicate names
func uniqueName(name string) (bool, error) {

	currentList, err := read()
	if err != nil {
		return false, err
	}

	for _, acc := range currentList {
		if name == acc.Name {
			return false, nil
		}
	}

	return true, nil
}

// getCurrentFromConfig retrieves the currently active/set API key from the
// config file
func getCurrentFromConfig() (string, error) {
	var currentKey string

	currentList, err := read()
	if err != nil {
		return currentKey, err
	}

	for i := range currentList {
		if currentList[i].Current == false {
			continue
		} else {
			currentKey = currentList[i].Key
		}
	}

	return currentKey, nil
}

func hasCurrent() (bool, error) {

	currentList, err := read()
	if err != nil {
		return false, err
	}

	for i := range currentList {
		if currentList[i].Current == false {
			continue
		} else {
			return true, nil
		}
	}

	return false, nil
}

// configPath returns the home directory of the current user
func configPath() (string, error) {
	x := os.Getenv("BINDPLANE_CONFIG_FILE")
	if len(x) > 0 {
		return x, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir + "/.bpcli", nil
}

// accountExists checks the config file to see whether a given account exists
func accountExists(name string) (bool, error) {

	currentList, err := read()
	if err != nil {
		return false, err
	}

	for i := range currentList {
		if name != currentList[i].Name {
			continue
		} else {
			return true, nil
		}
	}

	return false, nil
}

// checkConfig determines if a config file exists and whether it is empty
func checkConfig() (bool, error) {

	file, err := read()
	if err != nil {
		return false, err
	}

	if !(len(file) > 0) {
		return true, errors.New("The accounts list is empty")
	}

	return true, nil
}
