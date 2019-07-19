package sdk

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/BlueMedoraPublic/bpcli/util/uuid"
)

// SourceConfigGet type describes a source configuration
type SourceConfigGet struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	SourceType struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		URL    string `json:"url"`
		DocURL string `json:"doc_url"`
	} `json:"source_type"`
	CollectionInterval int          `json:"collection_interval"`
	CreatedAt          string       `json:"created_at"`
	Credentials        []Credential `json:"credentials"`
	Configuration      interface{}  `json:"configuration"`
	Status             string       `json:"status"`
	StatusReportedAt   string       `json:"status_reported_at"`
	StatusMessage      string       `json:"status_message"`
	Stopped            bool         `json:"stopped"`
	Collector          struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		URL        string `json:"url"`
		Version    string `json:"version"`
		Hostname   string `json:"hostname"`
		Status     string `json:"status"`
		NumSources int    `json:"num_sources"`
	} `json:"collector"`
}

// SourceConfigCreate type describes a source configuration to be
// created
type SourceConfigCreate struct {
	CollectionInterval int         `json:"collection_interval"`
	CollectorID        string      `json:"collector_id"`
	Configuration      interface{} `json:"configuration"`
	Credentials        struct {
		Credentials string `json:"credentials"`
	} `json:"credentials"`
	Name       string `json:"name"`
	SourceType string `json:"source_type"`
}

// SourceType type describes a source type configuration
type SourceType struct {
	ID                        string           `json:"id"`
	Name                      string           `json:"name"`
	URL                       string           `json:"url"`
	DocURL                    string           `json:"doc_url"`
	Version                   string           `json:"version"`
	DefaultCollectionInterval int              `json:"default_collection_interval"`
	CredentialTypes           []CredentialType `json:"credential_types"`
	ConnectionParameters      interface{}      `json:"connection_parameters"`
}

// SourceTypeTemplate type describes a source type template configuration
type SourceTypeTemplate struct {
	Name               string `json:"name"`
	SourceType         string `json:"source_type"`
	CollectorID        string `json:"collector_id"`
	CollectionInterval int    `json:"collection_interval"`
	Credentials        struct {
		Credentials string `json:"credentials"`
	} `json:"credentials"`
	Configuration interface{} `json:"configuration"`
}

// SourceCreateResponse type describes the json body returned by
// the source create endpoint
type SourceCreateResponse struct {
	JobID string `json:"job_id"`
	URL   string `json:"url"`
}

// GetSource will return a source object
func (bp BindPlane) GetSource(id string) (SourceConfigGet, error) {
	var s SourceConfigGet
	body, err := bp.APICall("get", bp.paths.sources+"/"+id, nil)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	return s, err
}

// GetSources will return a source object
func (bp BindPlane) GetSources() ([]SourceConfigGet, error) {
	var s []SourceConfigGet
	body, err := bp.APICall("get", bp.paths.sources, nil)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	return s, err
}

// GetSourceTemplate will return a SourceTypeTemplate with
// default values
func (bp BindPlane) GetSourceTemplate(id string) (SourceTypeTemplate, error) {
	var t SourceTypeTemplate
	body, err := bp.APICall("GET", bp.paths.sourceTypes+"/"+id+"/template", nil)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)
	return t, err
}

// GetSourceType will return a source type
func (bp BindPlane) GetSourceType(id string) (SourceType, error) {
	var t SourceType
	body, err := bp.APICall("get", bp.paths.sourceTypes+"/"+id, nil)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)
	return t, err
}

// ListSourceTypes will return an array of available source types
func (bp BindPlane) ListSourceTypes() ([]SourceType, error) {
	var t []SourceType
	body, err := bp.APICall("get", bp.paths.sourceTypes, nil)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)
	return t, err
}

// DeleteSource will delete a configured source
func (bp BindPlane) DeleteSource(id string) ([]byte, error) {
	return bp.APICall("delete", bp.paths.sources+"/"+id, nil)
}

// CreateSource will configure a new source, returning the
// http response body, http status code, and an error
func (bp BindPlane) CreateSource(payload []byte) (SourceCreateResponse, error) {
	var resp SourceCreateResponse
	body, err := bp.APICall("post", bp.paths.sources, payload)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	return resp, err
}

// Validate validates a SourceConfigCreate configuration
func (s SourceConfigCreate) Validate() error {
	var msg string
	if s.CollectionInterval < 1 {
		msg = msg + "\ncollection interval cannot be less than 1"
	}
	if !uuid.IsUUID(s.CollectorID) {
		msg = msg + "\ncollector id is invalid"
	}
	if !uuid.IsUUID(s.Credentials.Credentials) {
		msg = msg + "\ncredentials id is not a valid UUID"
	}
	if len(s.Name) == 0 {
		msg = msg + "\nname is not present"
	}
	if len(s.SourceType) == 0 {
		msg = msg + "\nsource type is not present"
	}

	if len(msg) > 0 {
		return errors.New("failed to validate source config:" + msg)
	}
	return nil
}

// Print prints a Source
func (s SourceConfigGet) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", s.ID, "source_type:", s.SourceType.ID, "status:", s.Status)
	return nil
}

// Print prints a SourceType
func (s SourceType) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", s.ID, "doc:", s.DocURL)
	return nil
}

// Print prints a SourceTypeTemplate
func (s SourceTypeTemplate) Print() error {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf(string(b))
	return nil
}

// Print prints a SourceCreateResponse
func (s SourceCreateResponse) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("job_id:", s.JobID, "doc_url:", s.URL)
	return nil
}
