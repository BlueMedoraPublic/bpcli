package sdk

import (
	"testing"
)

func testPopulatedLogSourceConfig() LogSourceConfig {
	x := newLogSourceConfig()
	x.ID = "abc"
	x.Name = "abc"
	x.Source.ID = "abc"
	x.Source.Name = "abc"
	x.Source.Version = "1.0"
	x.Configuration["ReadFromHead"] = true
	x.Configuration["ReadInterval"] = 2
	x.Configuration["MaxReads"] = "2"
	return x
}

func TestValidateCreateGood(t *testing.T) {
	x := testPopulatedLogSourceConfig()
	x.ID = ""
	if err := x.ValidateCreate(); err != nil {
		t.Errorf("Expected ValidateCreate() to return a nil error when ID is empty, got: " + err.Error())
	}
}

func TestValidateCreateBadID(t *testing.T) {
	x := testPopulatedLogSourceConfig()
	x.ID = "abc"
	if err := x.ValidateCreate(); err == nil {
		t.Errorf("Expected ValidateCreate() to return an error when ID is not set, got nil")
	}
}

func TestValidateCreateBadName(t *testing.T) {
	x := testPopulatedLogSourceConfig()
	x.ID = ""
	x.Name = ""
	if err := x.ValidateCreate(); err == nil {
		t.Errorf("Expected ValidateCreate() to return an error when Name is not set, got nil")
	}
}

func TestValidateCreateBadSourceID(t *testing.T) {
	x := testPopulatedLogSourceConfig()
	x.ID = ""
	x.Source.ID = ""
	if err := x.ValidateCreate(); err == nil {
		t.Errorf("Expected ValidateCreate() to return an error when Source.ID is not set, got nil")
	}
}

func TestValidateCreateBadSourceName(t *testing.T) {
	x := testPopulatedLogSourceConfig()
	x.ID = ""
	x.Source.Name = ""
	if err := x.ValidateCreate(); err == nil {
		t.Errorf("Expected ValidateCreate() to return an error when Source.Name is not set, got nil")
	}
}

func TestValidateCreateBadSourceVersion(t *testing.T) {
	x := testPopulatedLogSourceConfig()
	x.ID = ""
	x.Source.Version = ""
	if err := x.ValidateCreate(); err == nil {
		t.Errorf("Expected ValidateCreate() to return an error when Source.Version is not set, got nil")
	}
}

func TestValidateCreateBadConfiguration(t *testing.T) {
	x := testPopulatedLogSourceConfig()
	x.ID = ""
	x.Configuration = make(map[string]interface{})
	if err := x.ValidateCreate(); err == nil {
		t.Errorf("Expected ValidateCreate() to return an error when Configuration is an empty map, got nil")
	}
}
