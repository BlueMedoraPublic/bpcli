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

func TestGetLogSourceTypeParamsPath(t *testing.T) {
	x := GetLogSourceTypeParamsPath(apiVersion, "mysql")
	if x != (apiVersion + logSourceTypesPath + "/mysql/parameters") {
		t.Errorf("Expected GetLogSourceTypeParamsPath() to return " + apiVersion + logSourceTypesPath + "/mysql/parameters")
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

func TestGetLogAgentPath(t *testing.T) {
	x := GetLogAgentPath(apiVersion, "abc")
	if x != GetLogAgentsAllPath(apiVersion)+"/"+"abc" {
		t.Errorf("Expected GetLogAgentsAllPath() to return " + GetLogAgentsAllPath(apiVersion) + "/" + "abc")
	}
}

func TestGetLogAgentTaskDetailsPath(t *testing.T) {
	x := GetLogAgentTaskDetailsPath(apiVersion, "abc", "xyz")
	if x != GetLogAgentPath(apiVersion, "abc")+"/tasks/"+"xyz" {
		t.Errorf("Expected GetLogAgentTaskDetailsPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/tasks/" + "xyz")
	}
}

func TestGetLogAgentSourcesPath(t *testing.T) {
	x := GetLogAgentSourcesPath(apiVersion, "abc")
	if x != GetLogAgentPath(apiVersion, "abc")+"/sources" {
		t.Errorf("Expected GetLogAgentSourcesPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/sources")
	}
}

func TestGetLogAgentDestinationsPath(t *testing.T) {
	x := GetLogAgentDestinationsPath(apiVersion, "abc")
	if x != GetLogAgentPath(apiVersion, "abc")+"/destinations" {
		t.Errorf("Expected GetLogAgentDestinationsPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/destinations")
	}
}

func TestGetLogAgentUpdatePath(t *testing.T) {
	x := GetLogAgentUpdatePath(apiVersion, "abc")
	if x != GetLogAgentPath(apiVersion, "abc")+"/update_agent" {
		t.Errorf("Expected GetLogAgentUpdatePath() to return " + GetLogAgentPath(apiVersion, "abc") + "/update_agent")
	}
}

func TestGetLogAgentDeploySourceConfigPath(t *testing.T) {
	x := GetLogAgentDeploySourceConfigPath(apiVersion, "abc")
	if x != GetLogAgentPath(apiVersion, "abc")+"/deploy_source_config" {
		t.Errorf("Expected GetLogAgentDeploySourceConfigPath() to return" + GetLogAgentPath(apiVersion, "abc") + "/deploy_source_config")
	}
}

func TestGetLogAgentDeployDestinationConfigPath(t *testing.T) {
	x := GetLogAgentDeployDestinationConfigPath(apiVersion, "abc")
	if x != GetLogAgentPath(apiVersion, "abc")+"/deploy_destination_config" {
		t.Errorf("Expected GetLogAgentDeployDestinationConfigPath() to return " + GetLogAgentPath(apiVersion, "abc") + "/deploy_destination_config")
	}
}

func TestGetLogDestinatinationTypesPath(t *testing.T) {
	x := GetLogDestinationTypesPath(apiVersion)
	if x != apiVersion+logDestinationTypesPath {
		t.Errorf("Expected GetLogDestinationTypesPath() to return " + apiVersion + logDestinationTypesPath)
	}
}

func TestGetLogDestinationTypesParametersPath(t *testing.T) {
	x := GetLogDestinationTypesParametersPath(apiVersion, "abc")
	if x != GetLogDestinationTypesPath(apiVersion)+"/abc/parameters" {
		t.Errorf("Expected GetLogDestinationTypesParametersPath() to return " + GetLogDestinationTypesPath(apiVersion) + "abc")
	}
}

func TestGetLogDestinationConfigsAllPath(t *testing.T) {
	x := GetLogDestinationConfigsAllPath(apiVersion)
	if x != apiVersion+logDestinationConfigsPath {
		t.Errorf("Expected GetLogDestinationConfigsAllPath() to return " + apiVersion + logDestinationConfigsPath)
	}
}

func TestGetLogDestinationConfigPath(t *testing.T) {
	x := GetLogDestinationConfigPath(apiVersion, "abc")
	if x != GetLogDestinationConfigsAllPath(apiVersion)+"/abc" {
		t.Errorf("Expected GetLogDestinationConfigPath() to return " + GetLogDestinationConfigsAllPath(apiVersion) + "/abc")
	}
}

func TestGetLogDestinationUpdatePath(t *testing.T) {
	x := GetLogDestinationUpdatePath(apiVersion, "abc")
	if x != GetLogDestinationConfigsAllPath(apiVersion)+"/abc/update_destination" {
		t.Errorf("Expected GetLogDestinationUpdatePath() to return" + GetLogDestinationConfigsAllPath(apiVersion) + "/abc/update_destination")
	}
}
