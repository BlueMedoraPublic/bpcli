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
