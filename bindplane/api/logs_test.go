package api

import (
    "testing"
)

func TestGetLogsourcetypesPath(t *testing.T) {
    x := GetLogSourceTypesPath(apiVersion)
    if x != (apiVersion + logSourceTypesPath) {
        t.Errorf("Expected GetLogSourcetypesPath() to return " + apiVersion + logSourceTypesPath)
    }
}

func TestGetLogSourceTypeParamsPath(t *testing.T) {
    x := GetLogSourceTypeParamsPath(apiVersion, "mysql")
    if x != (apiVersion + logSourceTypesPath + "/mysql/parameters") {
        t.Errorf("Expected GetLogSourceTypeParamsPath() to return " + apiVersion + logSourceTypesPath + "/mysql/parameters")
    }
}

func TestGetLogSourceConfigsPath(t *testing.T) {
    x := GetLogSourceConfigsPath(apiVersion)
    if x != apiVersion + logSourceConfigsPath {
        t.Errorf("Expected GetLogSourceConfigsPath() to return " + apiVersion + logSourceConfigsPath)
    }
}

func TestGetLogAgentsAllPath(t *testing.T) {
    x := GetLogAgentsAllPath(apiVersion)
    if x != apiVersion + logAgentsPath {
        t.Errorf("Expected GetLogAgentsAllPath() to return " + apiVersion + logAgentsPath)
    }
}

func TestGetLogAgentInstallCommandPath(t *testing.T) {
    x := GetLogAgentInstallCommandPath(apiVersion)
    if x != GetLogAgentsAllPath(apiVersion) + "/install_commands" {
        t.Errorf("Expected GetLogAgentInstallCommandPath() to return " + GetLogAgentsAllPath(apiVersion) + "/install_commands")
    }
}

func TestGetLogAgentPath(t *testing.T) {
    x := GetLogAgentPath(apiVersion, "abc")
    if x != GetLogAgentsAllPath(apiVersion) + "/" + "abc" {
        t.Errorf("Expected GetLogAgentsAllPath() to return " + GetLogAgentsAllPath(apiVersion) + "/" + "abc")
    }
}

func TestGetLogAgentTaskDetailsPath(t *testing.T) {
    x := GetLogAgentTaskDetailsPath(apiVersion, "abc", "xyz")
    if x != GetLogAgentPath(apiVersion, "abc") + "/tasks/" + "xyz" {
        t.Errorf("Expected GetLogAgentTaskDetailsPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/tasks/" + "xyz")
    }
}

func TestGetLogAgentSourcesPath(t *testing.T) {
    x := GetLogAgentSourcesPath(apiVersion, "abc")
    if x != GetLogAgentPath(apiVersion, "abc") + "/sources" {
        t.Errorf("Expected GetLogAgentSourcesPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/sources")
    }
}

func TestGetLogAgentDestinationsPath(t *testing.T) {
    x := GetLogAgentDestinationsPath(apiVersion, "abc")
    if x != GetLogAgentPath(apiVersion, "abc") + "/destinations" {
        t.Errorf("Expected GetLogAgentDestinationsPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/destinations")
    }
}

func TestGetLogAgentUpdatePath(t *testing.T) {
    x := GetLogAgentUpdatePath(apiVersion, "abc")
    if x != GetLogAgentPath(apiVersion, "abc") + "/update_agent" {
        t.Errorf("Expected GetLogAgentUpdatePath() to return " + GetLogAgentPath(apiVersion, "abc") + "/update_agent")
    }
}

func TestGetLogAgentDeploySourceConfigPath(t *testing.T) {
    x := GetLogAgentDeploySourceConfigPath(apiVersion, "abc")
    if x != GetLogAgentPath(apiVersion, "abc") + "/deploy_source_config" {
        t.Errorf("Expected GetLogAgentDeploySourceConfigPath() to return" + GetLogAgentPath(apiVersion, "abc") + "/deploy_source_config")
    }
}

func TestGetLogAgentDeployDestinationConfigPath(t *testing.T) {
    x := GetLogAgentDeployDestinationConfigPath(apiVersion, "abc")
    if x != GetLogAgentPath(apiVersion, "abc") + "/deploy_destination_config" {
        t.Errorf("Expected GetLogAgentDeployDestinationConfigPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/deploy_destination_config")
    }
}
