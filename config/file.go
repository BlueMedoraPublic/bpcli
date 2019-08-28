package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/user"
)

// checkConfig determines if a config file exists and whether it is empty
func checkConfig() (bool, error) {
	accounts, err := read()
	if err != nil {
		return false, err
	}

	// return true with or without an errror
	if !(len(accounts) > 0) {
		return true, errors.New("The accounts list is empty")
	}
	return true, nil
}

// configPath returns the home directory of the current user
// if BINDPLANE_CONFIG_FILE is not set
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
	return ioutil.WriteFile(filePath, list, 0600)
}

// create creates the config file
func create() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	emptyFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer emptyFile.Close()
	return nil
}
