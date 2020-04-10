package config

import (
	"testing"
)

func TestUniqueUUID(t *testing.T) {
	// using fake UUIDs generated from https://www.uuidgenerator.net/
	u := "459371a9-73df-4a0b-b212-e141ef4e3136"
	a := []account{}
	a = append(a, account{Name: "bob", Key: u})

	if uniqueUUID(a, u) == true {
		t.Errorf("Expected uniqueUUID() to return false when passing in a UUID that is already in the account list")
	}
}

func TestUniqueName(t *testing.T) {
	// using fake UUIDs generated from https://www.uuidgenerator.net/
	u := "459371a9-73df-4a0b-b212-e141ef4e3136"
	a := []account{}
	a = append(a, account{Name: "bob", Key: u})

	if uniqueName(a, "bob") == true {
		t.Errorf("Expected uniqueUUID() to return false when passing in a UUID that is already in the account list")
	}
}
