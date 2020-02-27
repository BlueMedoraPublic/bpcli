package sdk

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/pkg/errors"
)

// LogAgent represents a BindPlane log agent
type LogAgent struct {
    ID            string `json:"id"`
    Name          string `json:"name"`
    Version       string `json:"version"`
    LatestVersion string `json:"latest_version"`
    Status        string `json:"status"`
}

// LogAgentSource represents a BindPlane log agent's source configuration
type LogAgentSource struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Version       string `json:"version"`
	LatestVersion string `json:"latest_version"`
	Type          string `json:"type"`
}

// LogAgentTask represents a BindPlane log agent
type LogAgentTask struct {
    ID      string `json:"id"`
	AgentID string `json:"agent_id"`
	Name    string `json:"name"`
	State   string `json:"state"`
}

// InstallCMDLogAgent returns the install commands for installing
// the bindplane log agent
func (bp BindPlane) InstallCMDLogAgent() ([]byte, error) {
    uri := bp.paths.logs.agentInstallCmd
    body, err := bp.APICall(http.MethodGet, uri, nil)
    return body, err
}

// GetLogAgent returns a log agent
func (bp BindPlane) GetLogAgent(id string) (LogAgent, error) {
    var a LogAgent
    uri := bp.paths.logs.agents+"/"+id
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return a, err
    }

    err = json.Unmarshal(body, &a)
    return a, err
}

// DeleteLogAgent deletes a log agent
func (bp BindPlane) DeleteLogAgent(id string) error {
    uri := bp.paths.logs.agents+"/"+id
    _, err := bp.APICall(http.MethodDelete, uri, nil)
    return err
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

// GetLogAgentTask returns a task for a given agent
func (bp BindPlane) GetLogAgentTask(agentID, taskID string) (LogAgentTask, error) {
    var t LogAgentTask
    uri := bp.paths.logs.agents+"/"+agentID+"/tasks/"+taskID
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return t, err
    }

    err = json.Unmarshal(body, &t)
    return t, err
}

// ListLogAgentSources returns a log agent
func (bp BindPlane) ListLogAgentSources(id string) ([]LogAgentSource, error) {
    var s []LogAgentSource
    uri := bp.paths.logs.agents+"/"+id+"/sources"
    body, err := bp.APICall(http.MethodGet, uri, nil)
    if err != nil {
        return s, err
    }

    err = json.Unmarshal(body, &s)
    return s, err
}

// GetLogAgentSources returns a log agent
func (bp BindPlane) GetLogAgentSources(agentID, sourceID string) (LogAgentSource, error) {
    s, err := bp.ListLogAgentSources(agentID)
    if err != nil {
        return LogAgentSource{}, err
    }

    for _, source := range s {
        if source.ID == sourceID {
            return source, nil
        }
    }

    return LogAgentSource{}, errors.New("source with id " + sourceID + " was not found when reading agent sources. agent_id: " + agentID)
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

// Print prints a LogAgent
func (t LogAgentTask) Print(j bool) error {
    if j == true {
		b, err := json.MarshalIndent(t, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", t.ID, "agent_id:", t.AgentID, "name:", t.Name, "state:", t.State)
	return nil
}

// Print prints a LogAgentSource
func (s LogAgentSource) Print(j bool) error {
    if j == true {
		b, err := json.MarshalIndent(s, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", s.ID, "name:", s.Name, "Version:", s.Version)
	return nil
}