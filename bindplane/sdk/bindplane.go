package sdk

import (
	"os"

	"github.com/BlueMedoraPublic/bpcli/bindplane/api"
	"github.com/BlueMedoraPublic/bpcli/config"
	"github.com/BlueMedoraPublic/bpcli/util/httpclient"
	"github.com/pkg/errors"
)

const bindplaneAPIEndpoint = "BINDPLANE_API_ENDPOINT"
const bindplaneAPIVersion = "BINDPLANE_API_VERSION"


// BindPlane type stores the global configuration
type BindPlane struct {
	BaseURL    string
	APIKey     string
	APIVersion string

	paths struct {
		collectors      string
		credentials     string
		credentialTypes string
		jobs            string
		sources         string
		sourceTypes     string
	}
}

// Init checks set configuration values and sets defaults
// if required
func (bp *BindPlane) Init() error {
	if err := bp.setBaseURL(); err != nil {
		return err
	}

	if err := bp.setAPIKey(); err != nil {
		return err
	}

	if err := bp.setAPIVersion(); err != nil {
		return err
	}

	if apiVersionIsValid(bp.APIVersion) == false {
		return errors.New("API Version " + bp.APIVersion + " is not valid.")
	}

	bp.paths.collectors = api.GetCollectorPath(bp.APIVersion)
	bp.paths.credentials = api.GetCredentialPath(bp.APIVersion)
	bp.paths.jobs = api.GetJobPath(bp.APIVersion)
	bp.paths.credentialTypes = api.GetCredentialTypePath(bp.APIVersion)
	bp.paths.sources = api.GetSourcePath(bp.APIVersion)
	bp.paths.sourceTypes = api.GetSourceTypePath(bp.APIVersion)

	return nil
}

/*
APICall takes a HTTP method, relative api path, and payload
returns the API response and an error
example: APICall("GET", "/v1/collectors", nil)
*/
func (bp BindPlane) APICall(method string, relativePath string, payload []byte) ([]byte, error) {
	return httpclient.Request(method, bp.BaseURL+relativePath, payload, bp.APIKey)
}

func (bp *BindPlane) setBaseURL() error {
	// if already set programmatically
	if len(bp.BaseURL) > 0 {
		return nil
	}

	// if env is set
	x := os.Getenv(bindplaneAPIEndpoint)
	if len(x) > 0 {
		bp.BaseURL = x
		return nil
	}

	// set default
	if len(bp.BaseURL) < 1 {
		bp.BaseURL = api.GetDefaultBaseURL()
	}
	return nil
}

func (bp *BindPlane) setAPIKey() error {
	// var apiKey string

	// Checks current API Key string length
	if len(bp.APIKey) == 0 {
		apiKey, err := config.CurrentAPIKey()
		if err != nil {
			return err
		}
		// Set API Key
		bp.APIKey = apiKey
	}

	return nil
}

func (bp *BindPlane) setAPIVersion() error {
	if len(bp.APIVersion) > 0 {
		return nil
	}

	x := os.Getenv(bindplaneAPIVersion)
	if len(x) > 0 {
		bp.APIVersion = x
		return nil
	}

	bp.APIVersion = api.GetDefaultVersion()
	return nil
}

func apiVersionIsValid(v string) bool {
	for _, version := range api.Versions() {
		if v == version {
			return true
		}
	}
	return false
}
