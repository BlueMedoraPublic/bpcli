package api

import (
    "testing"
)

var apiVersion = GetDefaultVersion()

func TestApiConst(t *testing.T) {
    if defaultBaseURL != "https://public-api.bindplane.bluemedora.com" {
        t.Errorf("Expected defaultBaseURL to be https://public-api.bindplane.bluemedora.com")
    }

    if defaultVersion != "/v1" {
        t.Errorf("Expected defaultVersion to be /v1")
    }
}

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
