package api

import (
	"testing"
)

func TestVersions(t *testing.T) {
	x := Versions()
	if len(x) != 1 {
		t.Errorf("Expected Versions() to return a single api version as only '/v1' has been implemented")
	}

	foundV1 := false
	for _, v := range x {
		if v == "/v1" {
			foundV1 = true
		}
	}
	if foundV1 != true {
		t.Errorf("Expected Versions() to include '/v1'")
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

/*

Begin api path testing using default api path

*/
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
