package sdk

import (
    "testing"
)

func TestValidateCreateGood(t *testing.T) {
    x := newLogSourceConfig()

    x.ID = "" // id should be empty when creating a source
    x.Name = "abc"
    x.Source.ID = "abc"
    x.Source.Name = "abc"
    x.Source.Version = "1.0"
    x.Configuration["ReadFromHead"] = true
    x.Configuration["ReadInterval"] = 2
    x.Configuration["MaxReads"] = "2"

    if err := x.ValidateCreate(); err != nil {
        t.Errorf("Expected ValidateCreate() to return a nil error, got: " + err.Error())
    }
}

func TestValidateCreateBadID(t *testing.T) {
    x := newLogSourceConfig()

    x.ID = "abc" // id should be empty when creating a source
    x.Name = "abc"
    x.Source.ID = "abc"
    x.Source.Name = "abc"
    x.Source.Version = "1.0"
    x.Configuration["ReadFromHead"] = true
    x.Configuration["ReadInterval"] = 2
    x.Configuration["MaxReads"] = "2"

    if err := x.ValidateCreate(); err == nil {
        t.Errorf("Expected ValidateCreate() to return an error when ID is set, got nil")
    }
}
