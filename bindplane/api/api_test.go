package api

import (
    "testing"
)

var apiVersion = GetDefaultVersion()

func TestGetDefaultBaseURL(t *testing.T) {
	x := GetDefaultBaseURL()
	if x != defaultBaseURL {
		t.Errorf("Expected GetDefaultBaseURL() to return '" + defaultBaseURL + "', got " + x)
	}
}

func TestGetDefaultVersion(t *testing.T) {
	x := GetDefaultVersion()
	if x != defaultVersion {
		t.Errorf("Expected GetDefaultVersion() to return '" + defaultVersion + "', got " + x)
	}
}
