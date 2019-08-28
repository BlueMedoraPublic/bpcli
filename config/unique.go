package config

import (
	"errors"

	"github.com/BlueMedoraPublic/bpcli/util/uuid"
)

// uniqueUUID checks the account list for duplicate UUIDs
func uniqueUUID(accounts []account, key string) (bool, error) {
	if uuid.IsUUID(key) == false {
		return false, errors.New("The API key given is not a valid UUID")
	}

	for _, a := range accounts {
		if a.Key == key {
			return false, nil
		}
	}
	return true, nil
}

// uniqueName checks the users account list for duplicate names
func uniqueName(accounts []account, name string) (bool, error) {
	for _, a := range accounts {
		if name == a.Name {
			return false, nil
		}
	}
	return true, nil
}
