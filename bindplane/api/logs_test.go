package api

import (
	"testing"
)

func TestLogConsts(t *testing.T) {
	if logSourceTypesPath != "/logs/source_types" {
		t.Errorf("Expected logSourcetypesPath const to equal /logs/source_types")
	}

	if logSourceConfigsPath != "/logs/source_configs" {
		t.Errorf("Expected logSourceConfigsPath const to equal /logs/source_configs")
	}

	if logAgentsPath != "/logs/agents" {
		t.Errorf("Expected logAgentsPath const to equal /logs/agents")
	}

	if logDestinationTypesPath != "/logs/destination_types" {
		t.Errorf("Expected logDestinationTypesPath const to equal /logs/destination_types")
	}

	if logDestinationConfigsPath != "/logs/destination_configs" {
		t.Errorf("Expected logDestinationConfigsPath const to equal /logs/destination_configs")
	}

	if logTemplatesPath != "/logs/templates" {
		t.Errorf("Expected logTemplatesPath const to equal /logs/templates")
	}

}

func TestGetLogsourcetypesPath(t *testing.T) {
	x := GetLogSourceTypesPath(apiVersion)
	if x != (apiVersion + logSourceTypesPath) {
		t.Errorf("Expected GetLogSourcetypesPath() to return " + apiVersion + logSourceTypesPath)
	}
}

func TestGetLogSourceConfigsPath(t *testing.T) {
	x := GetLogSourceConfigsPath(apiVersion)
	if x != apiVersion+logSourceConfigsPath {
		t.Errorf("Expected GetLogSourceConfigsPath() to return " + apiVersion + logSourceConfigsPath)
	}
}

func TestGetLogAgentsAllPath(t *testing.T) {
	x := GetLogAgentsAllPath(apiVersion)
	if x != apiVersion+logAgentsPath {
		t.Errorf("Expected GetLogAgentsAllPath() to return " + apiVersion + logAgentsPath)
	}
}

func TestGetLogAgentInstallCommandPath(t *testing.T) {
	x := GetLogAgentInstallCommandPath(apiVersion)
	if x != GetLogAgentsAllPath(apiVersion)+"/install_commands" {
		t.Errorf("Expected GetLogAgentInstallCommandPath() to return " + GetLogAgentsAllPath(apiVersion) + "/install_commands")
	}
}

func TestGetLogDestinatinationTypesPath(t *testing.T) {
	x := GetLogDestinationTypesPath(apiVersion)
	if x != apiVersion+logDestinationTypesPath {
		t.Errorf("Expected GetLogDestinationTypesPath() to return " + apiVersion + logDestinationTypesPath)
	}
}

func TestGetLogDestinationConfigsAllPath(t *testing.T) {
	x := GetLogDestinationConfigsAllPath(apiVersion)
	if x != apiVersion+logDestinationConfigsPath {
		t.Errorf("Expected GetLogDestinationConfigsAllPath() to return " + apiVersion + logDestinationConfigsPath)
	}
}
