package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const envAPIKey = "BINDPLANE_API_KEY"

// account stores users BindPlane account information
type account struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	Current bool   `json:"current"`
}

// AddAccount appends an account to the configuration file
func AddAccount(name string, key string) error {
	envWarning()

	accounts, err := read()
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") == false {
			return err
		}

		if err := create(); err != nil {
			return err
		}
	}

	accounts, err = read()
	if err != nil {
		return err
	}

	if err := validateNewAccount(accounts, name, key); err != nil {
		return err
	}

	a := account{Name: name, Key: key, Current: false}
	newList := append(accounts, a)

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
	apiKey, err := currentAPIKeyENV()
	if err != nil {
		return apiKey, err
	}

	if apiKey == "" {
		apiKey, err = currentAccount()
		if err != nil {
			return "", errors.Wrap(err, "could not read api key from environment or configuration file")
		}
	}

	return apiKey, nil
}

// ListAccounts prints a formatted list of users read from the configuration file
func ListAccounts() error {
	envWarning()

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

	fmt.Println("List of Accounts and API Keys. * Denotes Current Account")

	// Print the list in a formatted way
	for _, acc := range currentList {
		if acc.Current == true {
			fmt.Println("* "+acc.Name, acc.Key)
		} else {
			fmt.Println(acc.Name, acc.Key)
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
	envWarning()

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

func currentAPIKeyENV() (string, error) {
	a := os.Getenv(envAPIKey)

	if a != "" {
		if _, err := uuid.Parse(a); err != nil {
			return "", errors.Wrap(err, envAPIKey+" is set but is not a valid UUID")
		}
	}
	return a, nil
}

func currentAccount() (string, error) {
	accounts, err := read()
	if err != nil {
		return "", err
	}

	for _, a := range accounts {
		if a.Current {
			if _, err := uuid.Parse(a.Key); err != nil {
				return "", errors.Wrap(err, "Found current account in config, '"+a.Name+"', however, the API key is not a valid UUID")
			}
			return a.Key, nil
		}
	}
	return "", noCurrentAccountError()
}

func accountExists(name string) (bool, error) {
	currentList, err := read()
	if err != nil {
		return false, err
	}

	for i := range currentList {
		if name == currentList[i].Name {
			return true, nil
		}
	}
	return false, nil
}

func validateNewAccount(accounts []account, name string, key string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return errors.New("The name cannot be an empty string")
	}

	if _, err := uuid.Parse(key); err != nil {
		return errors.Wrap(err, "The API Key given is not a valid UUID: "+key)
	}

	if uniqueUUID(accounts, key) == false {
		return errors.New("The API Key given already exists within the config file")
	}

	if uniqueName(accounts, name) == false {
		return errors.New("The name given already exists within the config file")
	}

	return nil
}

func envWarning() {
	x := os.Getenv("BINDPLANE_API_KEY")
	if len(x) > 0 {
		fmt.Fprintf(os.Stderr, "WARNING: BINDPLANE_API_KEY is set and will take precidence over the configuration file\n")
	}
}
