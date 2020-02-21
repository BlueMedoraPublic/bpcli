package sdk

import (
    "fmt"
    "net/http"
    "encoding/json"
)

// LogSourceType type represents a log source type
type LogSourceType struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	DocURL  string `json:"doc_url"`
}

// ListLogSourceTypes returns all available log source types
func (bp BindPlane) ListLogSourceTypes() ([]LogSourceType, error) {
    var s []LogSourceType
    body, err := bp.APICall(http.MethodGet, bp.paths.logs.sourceTypes, nil)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(body, &s)
    return s, err
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
