package api

const collectorPath = "/collectors"
const credentialPath = "/credentials"
const credentialTypePath = "/credential_types"
const jobPath = "/jobs"
const sourcePath = "/sources"
const sourceTypePath = "/source_types"

// GetCollectorPath returns the api path for the collectors
// endpoint
func GetCollectorPath(version string) string {
	return version + collectorPath
}

// GetCredentialPath returns the api path for the credentials
// endpoint
func GetCredentialPath(version string) string {
	return version + credentialPath
}

// GetCredentialTypePath returns the api path for the credentials
// endpoint
func GetCredentialTypePath(version string) string {
	return version + credentialTypePath
}

// GetJobPath returns the api path for the jobs
// endpoint
func GetJobPath(version string) string {
	return version + jobPath
}

// GetSourcePath returns the api path for the sources
// endpoint
func GetSourcePath(version string) string {
	return version + sourcePath
}

// GetSourceTypePath returns the api path for the sources
// endpoint
func GetSourceTypePath(version string) string {
	return version + sourceTypePath
}
