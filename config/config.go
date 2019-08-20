package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"encoding/json"

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

// AddAccount appends an account to the configuration file
func AddAccount(name string, key string) error {

	currentList, err := read()
	if err != nil {
		// Write to standard error in order to make sure the user
		// can see that we are taking additional action
		os.Stderr.WriteString("ERROR: " + err.Error())

		path, err := configPath()
		if err != nil {
			return err
		}

		emptyFile, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// Write to standard error in order to make sure the user
		// can see that we are taking additional action
		os.Stderr.WriteString("Creating a new file at: " + path)

		emptyFile.Close()
	}

	currentList, err = read()
	if err != nil {
		return err
	}

	if len(strings.TrimSpace(name)) == 0 {
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

/*
CurrentAPIKey returns the API key found in the environment,
or the 'current' API key found in the credentials file if
the environment is not set
*/
func CurrentAPIKey() (string, error) {
	apiKey, found, err := currentAPIKeyENV()

	// return an error if env is found but malformed
	if found == true && err != nil {
		return "", err
	}

	// return api key if found
	if found == true && err == nil {
		return apiKey, nil
	}

	apiKey, e := getCurrentFromConfig()
	if e != nil {
		// return both ENV and File errors
		return "", errors.Wrap(err, e.Error())
	}
	return apiKey, nil
}

// ListAccounts prints a formatted list of users read from the configuration file
func ListAccounts() error {

	currentList, err := read()
	if err != nil {
		return errors.Wrap(err, fileNotFoundError().Error())
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	if len(currentList) == 0 {
		return errors.New(path + " is empty, add an account with 'bpcli account add'")
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

// Remove erases an account from the configuration file
func Remove(name string) error {

	currentList, err := read()
	if err != nil {
		return errors.Wrap(err, fileNotFoundError().Error())
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
		os.Stderr.WriteString("No names match the given input" +
			"Name Given: " + name)
		return nil
	}

	newListBytes, err := json.Marshal(newList)
	if err != nil {
		return err
	}

	return write(newListBytes)
}

// SetCurrent sets a chosen account to be the current account being worked in
func SetCurrent(name string) error {

	currentList, err := read()
	if err != nil {
		return errors.Wrap(err, fileNotFoundError().Error())
	}

	b, err := accountExists(name)
	if err != nil {
		return err
	}
	if b == false {
		return accountNotFoundError(name)
	}

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

// currentAPIKeyENV returns the API key, true, and nil if
// the API key is found in the environment and is a valid uuid
// returns false if the environment is empty
func currentAPIKeyENV() (string, bool, error) {
	a := os.Getenv("BINDPLANE_API_KEY")

	if len(strings.TrimSpace(a)) == 0 {
		return "", false, errors.New("ERROR: The BINDPLANE_API_KEY environment variable is not set")
	}

	if !uuid.IsUUID(a) {
		return "", true, errors.New("ERROR: The BINDPLANE_API_KEY environment variable is not a valid uuid")
	}

	return a, true, nil
}

// getCurrentFromConfig retrieves the currently active/set API key from the
// config file
func getCurrentFromConfig() (string, error) {
	b, err := hasCurrent()
	if err != nil {
		return "", err
	} else if b == false {
		return "", errors.New("ERROR: credential file does not have an account set, use 'bpcli account set'")
	}

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

			if len(currentKey) <= 0 {
				return currentKey, errors.New(currentList[i].Name + " does not have a" +
					" valid API Key set")
			}

			if !uuid.IsUUID(currentKey) {
				return currentKey, errors.New(currentKey + " for " +
					currentList[i].Name + " is not a valid UUID")
			}
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
