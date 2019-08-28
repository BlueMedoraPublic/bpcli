package config

import (
    "errors"
)

func fileNotFoundError() error {
    return errors.New("The credentials file has not been found. "+
        "Use the `bpcli account add` command to generate the file "+
        "and add an account to the list")
}

func accountNotFoundError(name string) error {
    return errors.New("The account '" + name + "' does not exist in the credentials file")
}

func noCurrentAccountError() error {
    return errors.New("The configuration file does not have an account " +
        "set as 'current', use 'bpcli account set' to select an account " +
        "or set the environment variable BINDPLANE_API_KEY to bypass the configuration file")
}
