package config

// uniqueUUID checks the account list for duplicate UUIDs
func uniqueUUID(accounts []account, key string) bool {
	for _, a := range accounts {
		if a.Key == key {
			return false
		}
	}
	return true
}

// uniqueName checks the users account list for duplicate names
func uniqueName(accounts []account, name string) bool {
	for _, a := range accounts {
		if name == a.Name {
			return false
		}
	}
	return true
}
