package sdk

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/pkg/errors"
)

// LogTemplate represents a bindplane logs template
type LogTemplate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AgentGroup string `json:"agent_group"`
}

// LogTemplateCreate represents the object used to create
// a bindplane logs template
type LogTemplateCreate struct {
	Name                string   `json:"name"`
	SourceConfigIds     []string `json:"source_config_ids"`
	DestinationConfigID string   `json:"destination_config_id"`
	AgentGroup          string   `json:"agent_group"`
}

// GetLogTemplate returns a log template
func (bp BindPlane) GetLogTemplate(id string) (LogTemplate, error) {
    t := LogTemplate{}
    uri := bp.paths.logs.templates+"/"+id
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return t, err
    }

    err = json.Unmarshal(body, &t)
    return t, err
}

// CreateLogTemplate returns a log template
func (bp BindPlane) CreateLogTemplate(t LogTemplateCreate) error {
    return errors.New("creating templates is not currently supported by the API")
    if err := t.Verify(); err != nil {
        return errors.Wrap(err, "cannot create new template")
    }

    payload, err := json.Marshal(t)
    if err != nil {
        return err
    }

    uri := bp.paths.logs.templates
    _, err = bp.APICall(http.MethodPost, uri, payload)
    return err
}

// UpdateLogTemplate returns a log template
func (bp BindPlane) UpdateLogTemplate(id string, t LogTemplateCreate) error {
    if err := t.Verify(); err != nil {
        return errors.Wrap(err, "cannot create new template")
    }

    payload, err := json.Marshal(t)
    if err != nil {
        return err
    }

    uri := bp.paths.logs.templates+"/"+id
    _, err = bp.APICall(http.MethodPut, uri, payload)
    return err
}

// ListLogTemplates returns all log templates
func (bp BindPlane) ListLogTemplates() ([]LogTemplate, error) {
    t := []LogTemplate{}
    uri := bp.paths.logs.templates
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return t, err
    }

    err = json.Unmarshal(body, &t)
    return t, err
}

// DeleteLogTemplate returns a log template
func (bp BindPlane) DeleteLogTemplate(id string) error {
    uri := bp.paths.logs.templates+"/"+id
    _, err := bp.APICall(http.MethodDelete, uri, nil)
    return err
}

// Print prints a LogTemplate
func (t LogTemplate) Print(j bool) error {
    if j == true {
		b, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", t.ID, "name:", t.Name, "agent_group:" + t.AgentGroup)
	return nil
}

// Verify validates a LogTemplateCreate type
func (t LogTemplateCreate) Verify() error {
    err := errors.New("log template validation failed")

    if t.Name == "" {
        return errors.Wrap(err, "name field is empty")
    }

    if len(t.SourceConfigIds) == 0 {
        return errors.Wrap(err, "source_config_ids is empty")
    }

    for i, sourceID := range t.SourceConfigIds {
        if sourceID == "" {
            p := strconv.Itoa(i)
            return errors.Wrap(err, "source config id in source_config_id["+p+"] list is empty")
        }
    }

    if t.DestinationConfigID == "" {
        return errors.Wrap(err, "destination_config_id is empty")
    }

    if t.AgentGroup == "" {
        return errors.Wrap(err, "agent_group is empty")
    }

    return nil
}
