package sdk

import (
	"testing"
)

/*
Ensure that SourceConfigGet uses []Credential rather than
a nested []struct, as some packages importing this library
may rely on this. Switching to a nested struct should be
considered a breaking change.
*/
func SourceConfigGetType(t *testing.T) {
	sourceConf := SourceConfigGet{}
	cred := Credential{ID: "myid"}
	sourceConf.Credentials = append(sourceConf.Credentials, cred)
	if len(sourceConf.Credentials) != 1 {
		t.Errorf("Expected sdk.SourceConfGet.Credentials to contain one credential object")
	}
}

func TestValidate(t *testing.T) {
	var s SourceConfigCreate

	// test valid config
	s = getValidSourceConfigCreate()
	if s.Validate() != nil {
		t.Errorf("Expected Validate() to not return an error when using a valid SourceConfigCreate struct\n" +
			"**VALIDATION ERROR MESSAGE**\n" +
			s.Validate().Error())
	}

	// test collection interval
	s = getValidSourceConfigCreate()
	s.CollectionInterval = -1
	if s.Validate() == nil {
		t.Errorf("Expected Validate() to return an error when using an invalid collection interval")
	}

	// test collector id
	s = getValidSourceConfigCreate()
	s.CollectorID = ""
	if s.Validate() == nil {
		t.Errorf("Expected Validate() to return an error when using an empty collector id")
	}

	// test credentials
	s = getValidSourceConfigCreate()
	s.Credentials.Credentials = ""
	if s.Validate() != nil {
		t.Errorf("Expected Validate() to a nil error when using an empty credential" +
			"**VALIDATION ERROR MESSAGE**\n" +
			s.Validate().Error())
	}

	// test name
	s = getValidSourceConfigCreate()
	s.Name = ""
	if s.Validate() == nil {
		t.Errorf("Expected Validate() to return an error when using an empty name")
	}

	// test source type
	s = getValidSourceConfigCreate()
	s.SourceType = ""
	if s.Validate() == nil {
		t.Errorf("Expected Validate() to return an error when using an empty source type")
	}
}

func getValidSourceConfigCreate() SourceConfigCreate {
	var s SourceConfigCreate
	s.CollectionInterval = 2
	s.CollectorID = "abcdefAB-0123-4ABC-ab12-CDEF01234567"
	s.Credentials.Credentials = "abcdefAB-0123-4ABC-ab12-CDEF01234567"
	s.Name = "abc"
	s.SourceType = "abc"
	return s
}
