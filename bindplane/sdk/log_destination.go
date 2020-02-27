package sdk

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/pkg/errors"
)

// LogDestType represents a logging destination
type LogDestType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	DocURL  string `json:"doc_url"`
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
    uri := bp.paths.logs.destTypes+"/"+id+"/parameters"
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
