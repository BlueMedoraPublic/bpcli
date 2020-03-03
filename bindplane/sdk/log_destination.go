package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// LogDestType represents a logging destination
type LogDestType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	DocURL  string `json:"doc_url"`
}

// LogDestConfig represents a logging destination configuration
type LogDestConfig struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Destination struct {
		ID      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"destination,omitempty"`
	Configuration map[string]interface{} `json:"configuraton,omitempty"`
	DestinationTypeID string `json:"destination_type_id,omitempty"`
	DestinationVersion string `json:"destination_version,omitempty"`
}

// GetLogDestType returns a logging destination type
func (bp BindPlane) GetLogDestType(id string) (LogDestType, error) {
	d, err := bp.ListLogDestTypes()
	if err != nil {
		return LogDestType{}, errors.Wrap(err, "cannot get destination type by id")
	}

	for _, d := range d {
		if d.ID == id {
			return d, nil
		}
	}

	err = errors.New("log destination type with id: " + id + " does not exist")
	return LogDestType{}, err
}

// GetLogDestTypeParameters returns a logging destination's configuraton
// parameters
func (bp BindPlane) GetLogDestTypeParameters(id string) ([]byte, error) {
	uri := bp.paths.logs.destTypes + "/" + id + "/parameters"
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	return body, err
}

// ListLogDestTypes lists available log destination types
func (bp BindPlane) ListLogDestTypes() ([]LogDestType, error) {
	var d []LogDestType
	uri := bp.paths.logs.destTypes
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// CreateLogDestConfig creates a destination config
func (bp BindPlane) CreateLogDestConfig(config LogDestConfig) (LogDestConfig, error) {
	var d LogDestConfig

	p, err := json.Marshal(config)
	if err != nil {
		return d, err
	}

	uri := bp.paths.logs.destConfigs
	body, err := bp.APICall(http.MethodPost, uri, p)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// CreateLogDestConfigRaw creates a destination config
func (bp BindPlane) CreateLogDestConfigRaw(config []byte) (LogDestConfig, error) {
	var d LogDestConfig
	uri := bp.paths.logs.destConfigs
	body, err := bp.APICall(http.MethodPost, uri, config)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// GetLogDestConfig returns a log destination config
func (bp BindPlane) GetLogDestConfig(id string) (LogDestConfig, error) {
	var d LogDestConfig
	uri := bp.paths.logs.destConfigs + "/" + id
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// UpdateLogDestConfig updates the log destination config version
// to the latest
func (bp BindPlane) UpdateLogDestConfig(id string) (LogDestConfig, error) {
	var d LogDestConfig
	uri := bp.paths.logs.destConfigs + "/" + id + "/update_destination_version"
	body, err := bp.APICall(http.MethodPut, uri, nil)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// DelLogDestConfig deletes a log destination config
func (bp BindPlane) DelLogDestConfig(id string) error {
	uri := bp.paths.logs.destConfigs + "/" + id
	_, err := bp.APICall(http.MethodDelete, uri, nil)
	return err
}

// ListLogDestConfigs lists available log destination types
func (bp BindPlane) ListLogDestConfigs() ([]LogDestConfig, error) {
	var d []LogDestConfig
	uri := bp.paths.logs.destConfigs
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// Print prints a LogDestType type
func (d LogDestType) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", d.ID, "name:", d.Name, "version:", d.Version, "doc_url:", d.DocURL)
	return nil
}

// Print print a LogDestConfig type
func (d LogDestConfig) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", d.ID, "name:", d.Name)
	return nil
}
