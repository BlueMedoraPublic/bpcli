package api

import (
	"testing"
)

var apiVersion = GetDefaultVersion()

func TestGetCollectorPath(t *testing.T) {
	x := GetCollectorPath(apiVersion)
	if x != (apiVersion + "/collectors") {
		t.Errorf("Expected GetCollectorPath() to return '" + apiVersion + "/collectors" + "', got " + x)
	}
}

func TestGetCredentialPath(t *testing.T) {
	x := GetCredentialPath(apiVersion)
	if x != (apiVersion + "/credentials") {
		t.Errorf("Expected GetCredentialPath() to return '" + apiVersion + "/credentials" + "', got " + x)
	}
}

func TestGetCredentialTypePath(t *testing.T) {
	x := GetCredentialTypePath(apiVersion)
	if x != (apiVersion + "/credential_types") {
		t.Errorf("Expected GetCredentialTypePath() to return '" + apiVersion + "/credential_types" + "', got " + x)
	}
}

func TestGetJobPath(t *testing.T) {
	x := GetJobPath(apiVersion)
	if x != (apiVersion + "/jobs") {
		t.Errorf("Expected GetJobPath() to return '" + apiVersion + "/jobs" + "', got " + x)
	}
}

func TestGetSourcePath(t *testing.T) {
	x := GetSourcePath(apiVersion)
	if x != (apiVersion + "/sources") {
		t.Errorf("Expected GetSourcePath() to return '" + apiVersion + "/sources" + "', got " + x)
	}
}

func TestGetSourceTypePath(t *testing.T) {
	x := GetSourceTypePath(apiVersion)
	if x != (apiVersion + "/source_types") {
		t.Errorf("Expected GetSourceTypePath() to return '" + apiVersion + "/source_types" + "', got " + x)
	}
}
