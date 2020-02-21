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

// GetLogDestinationTypesPath returns the api path for
// log listing log destinations
func GetLogDestinationTypesPath(version string) string {
	return version + logDestinationTypesPath
}

// GetLogDestinationConfigsAllPath returns the api path for
// listing destination configs
func GetLogDestinationConfigsAllPath(version string) string {
	return version + logDestinationConfigsPath
}

// GetLogTemplatesAllPath returns the api path for listing
// all templates
func GetLogTemplatesAllPath(version string) string {
	return version + logTemplatesPath
}
