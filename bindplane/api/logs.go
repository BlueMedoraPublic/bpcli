package api

const logSourceTypesPath = "/logs/source_types"
const logSourceConfigsPath = "/logs/source_configs"
const logAgentsPath = "/logs/agents"
const logDestinationTypesPath = "/logs/destination_types"
const logDestinationConfigsPath = "/logs/destination_configs"
const logTemplatesPath = "/logs/templates"

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

// GetLogAgentsAllPath returns the api path for all log agents
// endpoint
func GetLogAgentsAllPath(version string) string {
    return version + logAgentsPath
}

// GetLogAgentInstallCommandPath returns the api path for
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

// GetLogDestinationTypesPath returns the api path for
// log listing log destinations
func GetLogDestinationTypesPath(version string) string {
    return version + logDestinationTypesPath
}

// GetLogDestinationTypesParametersPath returns the api path for
// getting a destination type's parameters
func GetLogDestinationTypesParametersPath(version, destID string) string {
    return GetLogDestinationTypesPath(version) + "/" + destID + "/parameters"
}

// GetLogDestinationConfigsAllPath returns the api path for
// listing destination configs
func GetLogDestinationConfigsAllPath(version string) string {
    return version + logDestinationConfigsPath
}

// GetLogDestinationConfigPath returns the api path for
// getting the config for a given destination
func GetLogDestinationConfigPath(version, destID string) string {
    return GetLogDestinationConfigsAllPath(version) + "/" + destID
}

// GetLogDestinationUpdatePath returns the api path for
// updating an existing destination
func GetLogDestinationUpdatePath(version, destID string) string {
    return GetLogDestinationConfigsAllPath(version) + "/" + destID + "/update_destination"
}

// GetLogTemplatesAllPath returns the api path for listing
// all templates
func GetLogTemplatesAllPath(version string) string {
    return version + logTemplatesPath
}

// GetLogTemplatePath returns the api path for a specific
// template
func GetLogTemplatePath(version, templateID string) string {
    return version + logTemplatesPath + "/" + templateID
}
