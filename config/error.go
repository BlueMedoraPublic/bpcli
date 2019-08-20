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
