package api

const defaultBaseURL = "https://public-api.bindplane.bluemedora.com"
const defaultVersion = "/v1"

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
