package config

import (
    "os"
    "os/user"
)

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
