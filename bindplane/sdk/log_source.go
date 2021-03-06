package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		ID      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"source,omitempty"`

	// configuration is only returned when getting a specific
	// config, not during a list operation, therefore we omit
	// it when it is not present
	Configuration  map[string]interface{} `json:"configuration,omitempty"`
	CustomTemplate string                 `json:"custom_template,omitempty"`
	SourceTypeID   string                 `json:"source_type_id,omitempty"`
	SourceVersion  string                 `json:"source_version,omitempty"`
	CreatedAt      int64                  `json:"created_at,omitempty"`
	UpdatedAt      int64                  `json:"updated_at,omitempty"`
}

// GetLogSourceType returns a source type
func (bp BindPlane) GetLogSourceType(id string) (LogSourceType, error) {
	var s LogSourceType
	uri := bp.paths.logs.sourceTypes + "/" + id
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	return s, err
}

// GetLogSourceTypeParameters returns a source type template
func (bp BindPlane) GetLogSourceTypeParameters(id string) ([]byte, error) {
	uri := bp.paths.logs.sourceTypes + "/" + id + "/parameters"
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return body, err
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
	uri := bp.paths.logs.sourceConfigs + "/" + id
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

// CreateLogSourceConfig creates a log source config
func (bp BindPlane) CreateLogSourceConfig(config LogSourceConfig) (LogSourceConfig, error) {
	c := newLogSourceConfig()

	payload, err := json.Marshal(config)
	if err != nil {
		return c, err
	}

	uri := bp.paths.logs.sourceConfigs
	body, err := bp.APICall(http.MethodPost, uri, payload)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// CreateLogSourceConfigRaw creates a log source config
func (bp BindPlane) CreateLogSourceConfigRaw(config []byte) (LogSourceConfig, error) {
	c := newLogSourceConfig()
	uri := bp.paths.logs.sourceConfigs
	body, err := bp.APICall(http.MethodPost, uri, config)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// DeleteLogSourceConfig deletes a log source config
func (bp BindPlane) DeleteLogSourceConfig(id string) error {
	uri := bp.paths.logs.sourceConfigs + "/" + id
	_, err := bp.APICall(http.MethodDelete, uri, nil)
	return err
}

// UpdateVersionLogSourceConfig updates the log source config version
func (bp BindPlane) UpdateVersionLogSourceConfig(id string) (LogSourceConfig, error) {
	var c LogSourceConfig
	uri := bp.paths.logs.sourceConfigs + "/id" + id + "/update_source"
	body, err := bp.APICall(http.MethodPut, uri, nil)
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

// newLogSourceConfig returns a LogSourceConfig struct
// with an initilized Configuraton map. This helps prevent
// runtime nil points panics
func newLogSourceConfig() LogSourceConfig {
	return LogSourceConfig{
		Configuration: make(map[string]interface{}),
	}
}
