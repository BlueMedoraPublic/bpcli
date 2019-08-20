package config

import (
    "errors"

    "github.com/BlueMedoraPublic/bpcli/util/uuid"
)

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
