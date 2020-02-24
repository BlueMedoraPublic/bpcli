package sdk

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/pkg/errors"
)

// LogSourceType type represents a log source type
type LogSourceType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	DocURL  string `json:"doc_url"`
}

// LogSourceConfig type represents a log source config
type LogSourceConfig struct {
    // ID is always returned from the API but it is
    // not required when creating a new config
	ID     string `json:"id,omitempty"`
	Name   string `json:"name"`
	Source struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"source"`

    // configuration is only returned when getting a specific
    // config, not during a list operation, therefore we omit
    // it when it is not present
    Configuration map[string]interface{} `json:"configuration,omitempty"`
}

// GetLogSourceType returns a source type
func (bp BindPlane) GetLogSourceType(id string) (LogSourceType, error) {
    var s LogSourceType
    uri := bp.paths.logs.sourceTypes+"/"+id
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return s, err
    }

    err = json.Unmarshal(body, &s)
    return s, err
}

// GetLogSourceTypeParameters returns a source type template
// TODO: Finish once implemented in API
func (bp BindPlane) GetLogSourceTypeParameters(id string) ([]byte, error) {
    //var t LogSourceTypeTemplate
    uri := bp.paths.logs.sourceTypes+"/"+id+"/parameters"
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        //return t, err
        return nil, err
    }
    return body, err
    //err = json.Unmarshal(body, &t)
    //return s, err
}

// ListLogSourceTypes returns all available log source types
func (bp BindPlane) ListLogSourceTypes() ([]LogSourceType, error) {
    var s []LogSourceType
    uri := bp.paths.logs.sourceTypes
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return s, err
    }

    err = json.Unmarshal(body, &s)
    return s, err
}

// GetLogSourceConfig returns a specified log source config
func (bp BindPlane) GetLogSourceConfig(id string) (LogSourceConfig, error) {
    var c LogSourceConfig
    uri := bp.paths.logs.sourceConfigs+"/"+id
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return c, err
    }

    err = json.Unmarshal(body, &c)
    return c, err
}

// ListLogSourceConfigs returns all configured log sources
func (bp BindPlane) ListLogSourceConfigs() ([]LogSourceConfig, error) {
    var c []LogSourceConfig
    uri := bp.paths.logs.sourceConfigs
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return c, err
    }

    err = json.Unmarshal(body, &c)
    return c, err
}

// Print prints a LogSourceType
func (s LogSourceType) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", s.ID, "name:", s.Name, "version:", s.Version, "doc:", s.DocURL)
	return nil
}

// Print prints a LogSourceConfig
func (c LogSourceConfig) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(c, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", c.ID, "name:", c.Name, "source:", c.Source.ID)
	return nil
}

// ValidateCreate validates a LogSourceConfig to ensure
// it is formatted properly for creating a new config
func (c LogSourceConfig) ValidateCreate() error {
    if c.ID != "" {
        return errors.New("log source config id should not be set")
    }

    if c.Name == "" {
        return errors.New("log source config name is not set")
    }

    if c.Source.ID == "" {
        return errors.New("log source config source id is not set")
    }

    if c.Source.Name == "" {
        return errors.New("log source config source name is not set")
    }

    if c.Source.Version == "" {
        return errors.New("log source config source version is not set")
    }

    if len(c.Configuration) == 0 {
        return errors.New("log source config configuration is not set")
    }

    return nil
}
