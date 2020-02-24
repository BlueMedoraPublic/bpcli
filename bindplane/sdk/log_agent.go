package sdk

import (
    "fmt"
    "net/http"
    "encoding/json"
)

// LogAgent represents a BindPlane log agent
type LogAgent struct {
    ID            string `json:"id"`
    Name          string `json:"name"`
    Version       string `json:"version"`
    LatestVersion string `json:"latest_version"`
    Status        string `json:"status"`
}

// ListLogAgents returns all log agents
func (bp BindPlane) ListLogAgents() ([]LogAgent, error) {
    var a []LogAgent
    uri := bp.paths.logs.agents
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return a, err
    }

    err = json.Unmarshal(body, &a)
    return a, err
}

// Print prints a LogAgent
func (a LogAgent) Print(j bool) error {
    if j == true {
		b, err := json.MarshalIndent(a, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", a.ID, "name:", a.Name, "version:", a.Version, "status:", a.Status)
	return nil
}
