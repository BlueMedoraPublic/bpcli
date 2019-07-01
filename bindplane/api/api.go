package api

const defaultBaseURL     = "https://public-api.bindplane.bluemedora.com"
const defaultVersion  = "/v1"
const collectorPath      = "/collectors"
const credentialPath     = "/credentials"
const credentialTypePath = "/credential_types"
const jobPath            = "/jobs"
const sourcePath         = "/sources"
const sourceTypePath     = "/source_types"

// Versions returns a slice of valid api versions
func Versions() []string {
    return []string{"/v1"}
}

// GetDefaultBaseURL returns the default base url
func GetDefaultBaseURL() string {
    return defaultBaseURL
}

// GetDefaultVersion returns the default api version
func GetDefaultVersion() string {
    return defaultVersion
}

// GetCollectorPath returns the api path for the collectors
// endpoint
func GetCollectorPath(version string) string {
    return  version + collectorPath
}

// GetCredentialPath returns the api path for the credentials
// endpoint
func GetCredentialPath(version string) string {
    return  version + credentialPath
}

// GetCredentialTypePath returns the api path for the credentials
// endpoint
func GetCredentialTypePath(version string) string {
    return  version + credentialTypePath
}

// GetJobPath returns the api path for the jobs
// endpoint
func GetJobPath(version string) string {
    return  version + jobPath
}

// GetSourcePath returns the api path for the sources
// endpoint
func GetSourcePath(version string) string {
    return  version + sourcePath
}

// GetSourceTypePath returns the api path for the sources
// endpoint
func GetSourceTypePath(version string) string {
    return  version + sourceTypePath
}
