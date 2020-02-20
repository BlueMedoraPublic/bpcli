package api

const logSourceTypesPath = "/logs/source_types"
const logSourceConfigsPath = "/logs/source_configs"
const logAgentsPath = "/logs/agents"

// GetLogSourceTypesPath returns the api path for the log source
// types endpoint
func GetLogSourceTypesPath(version string) string {
    return version + logSourceTypesPath
}

// GetLogSourceTypeParamsPath returns the api path for a specific
// log source's parameters endpoint
func GetLogSourceTypeParamsPath(version, sourceID string) string {
    return GetLogSourceTypesPath(version) + "/" + sourceID + "/parameters"
}

// GetLogSourceConfigsPath returns the api path for all log
// source configs endpoint
func GetLogSourceConfigsPath(version string) string {
    return version + logSourceConfigsPath
}

// GetLogAgentsPath returns the api path for all log agents
// endpoint
func GetLogAgentsAllPath(version string) string {
    return version + logAgentsPath
}

// GetLogAgentsInstallCommandPath returns the api path for
// listing install commands endpoint
func GetLogAgentInstallCommandPath(version string) string {
    return GetLogAgentsAllPath(version) + "/install_commands"
}

// GetLogAgentPath returns the api path for a specific agent
func GetLogAgentPath(version, agentID string) string {
    return GetLogAgentsAllPath(version) + "/" + agentID
}

// GetLogAgentTaskDetailsPath returns the api path for
// getting a log agent task
func GetLogAgentTaskDetailsPath(version, agentID, taskID string) string {
    return GetLogAgentPath(version, agentID) + "/tasks/" + taskID
}

// GetLogAgentSourcesPath returns the api path for describing
// a given log agent's sources
func GetLogAgentSourcesPath(version, agentID string) string {
    return GetLogAgentPath(version, agentID) + "/sources"
}

// GetLogAgentDestinationsPath returns the api path for describing
// a given log agent's destinations
func GetLogAgentDestinationsPath(version, agentID string) string {
    return GetLogAgentPath(version, agentID) + "/destinations"
}

// GetLogAgentUpdatePath returns the api path for updating
// a given log agent's version
func GetLogAgentUpdatePath(version, agentID string) string {
    return GetLogAgentPath(version, agentID) + "/update_agent"
}

// GetLogAgentDeploySourceConfigPath returns the api path for
// deploying a source config ot an agent
func GetLogAgentDeploySourceConfigPath(version, agentID string) string {
    return GetLogAgentPath(version, agentID) + "/deploy_source_config"
}

// GetLogAgentDeployDestinationConfigPath returns the api path for
// deploying a source config ot an agent
func GetLogAgentDeployDestinationConfigPath(version, agentID string) string {
    return GetLogAgentPath(version, agentID) + "/deploy_destination_config"
}
