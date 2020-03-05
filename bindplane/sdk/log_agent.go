package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// LogAgentUpdateResp represents the object returned by
// the api when an update command is issued
type LogAgentUpdateResp struct {
	AgentID string `json:"agent_id"`
	TaskID  string `json:"task_id"`
}

// LogAgentSource represents a BindPlane log agent's source configuration
type LogAgentSource struct {
	SourceConfigID string `json:"source_config_id"`
	Name           string `json:"name"`
	Version        string `json:"version"`
	LatestVersion  string `json:"latest_version"`
	TypeID         string `json:"type_id"`
}

// LogAgentTask represents a BindPlane log agent task
type LogAgentTask struct {
	ID      string `json:"id"`
	AgentID string `json:"agent_id"`
	Name    string `json:"name"`
	State   string `json:"state"`
}

// LogAgentDest represents a Bindplane Log agent destination
// config
type LogAgentDest struct {
	DestinationConfigID string `json:"destination_config_id"`
	Name                string `json:"name"`
	Version             string `json:"version"`
	LatestVersion       string `json:"latest_version"`
	TypeID              string `json:"type_id"`
}

// InstallCMDLogAgent returns the install commands for installing
// the bindplane log agent
func (bp BindPlane) InstallCMDLogAgent(logAgentPlatform, templateID string) (string, error) {
	uri := bp.paths.logs.agentInstallCmd
	if templateID != "" {
		uri = uri + "?template_id=" + templateID
	}
	body, err := bp.APICall(http.MethodGet, uri, nil)

	platforms := make(map[string]string)
	if err := json.Unmarshal(body, &platforms); err != nil {
		return "", err
	}

	if logAgentPlatform == "all" {
		return string(body), err
	}

	p := []string{}
	for platform, command := range platforms {
		p = append(p, platform)
		if platform == logAgentPlatform {
			return command, nil
		}
	}

	// exit early if no platforms are returned
	if len(p) < 1 {
		err := errors.New("unexpected response from server, no install commands returned")
		return "", err
	}

	// safe to index p[0] because we know the slice is at
	// least length 1 from the check above ^
	valid := p[0]
	for i, p := range p {
		if i == 0 {
			continue
		}
		valid = valid + ", " + p
	}

	err = errors.New("platform is not supported: " + logAgentPlatform)
	return "", errors.Wrap(err, "supported platforms: "+valid)
}

// GetLogAgent returns a log agent
func (bp BindPlane) GetLogAgent(id string) (LogAgent, error) {
	var a LogAgent
	uri := bp.paths.logs.agents + "/" + id
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return a, err
	}

	err = json.Unmarshal(body, &a)
	return a, err
}

// UpdateLogAgent returns a log agent
func (bp BindPlane) UpdateLogAgent(id string) (LogAgentUpdateResp, error) {
	var a LogAgentUpdateResp
	uri := bp.paths.logs.agents + "/" + id + "/update_agent_version"
	body, err := bp.APICall(http.MethodPatch, uri, nil)
	if err != nil {
		return a, err
	}

	err = json.Unmarshal(body, &a)
	return a, err
}

// DeleteLogAgent deletes a log agent
func (bp BindPlane) DeleteLogAgent(id string) error {
	uri := bp.paths.logs.agents + "/" + id
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
	uri := bp.paths.logs.agents + "/" + agentID + "/tasks/" + taskID
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return t, err
	}

	err = json.Unmarshal(body, &t)
	return t, err
}

// ListLogAgentSources returns all log agents
func (bp BindPlane) ListLogAgentSources(id string) ([]LogAgentSource, error) {
	var s []LogAgentSource
	uri := bp.paths.logs.agents + "/" + id + "/sources"
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(body, &s)
	return s, err
}

// GetLogAgentSource returns a log agent's source by id
func (bp BindPlane) GetLogAgentSource(agentID, sourceID string) (LogAgentSource, error) {
	s, err := bp.ListLogAgentSources(agentID)
	if err != nil {
		return LogAgentSource{}, err
	}

	for _, source := range s {
		if source.SourceConfigID == sourceID {
			return source, nil
		}
	}

	err = errors.New("source with id " + sourceID + " was not found when reading agent sources. agent_id: " + agentID)
	return LogAgentSource{}, err
}

// DeployLogAgentSource deploys a source config to a log agent
func (bp BindPlane) DeployLogAgentSource(agentID, configID string) (LogAgentSource, error) {
	var d LogAgentSource
	payload := []byte("{\"source_config_id\":\"" + configID + "\"}")
	uri := bp.paths.logs.agents + "/" + agentID + "/deploy_source_config"
	body, err := bp.APICall(http.MethodPost, uri, payload)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// DeleteLogAgentSource deletes a source config from a log agent
func (bp BindPlane) DeleteLogAgentSource(agentID, sourceID string) error {
	uri := bp.paths.logs.agents + "/" + agentID + "/sources/" + sourceID
	_, err := bp.APICall(http.MethodDelete, uri, nil)
	return err
}

// ListLogAgentDest returns a log agent
func (bp BindPlane) ListLogAgentDest(id string) ([]LogAgentDest, error) {
	var d []LogAgentDest
	uri := bp.paths.logs.agents + "/" + id + "/destinations"
	body, err := bp.APICall(http.MethodGet, uri, nil)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// GetLogAgentDest returns a log agent's destination by id
func (bp BindPlane) GetLogAgentDest(agentID, destID string) (LogAgentDest, error) {
	d, err := bp.ListLogAgentDest(agentID)
	if err != nil {
		return LogAgentDest{}, err
	}

	for _, dest := range d {
		if dest.DestinationConfigID == destID {
			return dest, nil
		}
	}

	err = errors.New("destination with id " + destID + " was not found when reading agent sources. agent_id: " + agentID)
	return LogAgentDest{}, err
}

// DeployLogAgentDest deploys a destination config to a log agent
func (bp BindPlane) DeployLogAgentDest(agentID, configID string) (LogAgentDest, error) {
	var d LogAgentDest
	payload := []byte("{\"destination_config_id\":\"" + configID + "\"}")
	uri := bp.paths.logs.agents + "/" + agentID + "/deploy_destination_config"
	body, err := bp.APICall(http.MethodPost, uri, payload)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(body, &d)
	return d, err
}

// DeleteLogAgentDest deletes a destination config from a log agent
func (bp BindPlane) DeleteLogAgentDest(agentID, destID string) error {
	uri := bp.paths.logs.agents + "/" + agentID + "/destinations/" + destID
	_, err := bp.APICall(http.MethodDelete, uri, nil)
	return err
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

	fmt.Println("id:", s.SourceConfigID, "name:", s.Name, "version:", s.Version)
	return nil
}

// Print prints a LogAgentDest
func (d LogAgentDest) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(d, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("destination_config_id:", d.DestinationConfigID,
		"name:", d.Name,
		"version:", d.Version,
		"type_id:", d.TypeID)
	return nil
}

// Print prints a LogAgentUpdateResp
func (a LogAgentUpdateResp) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(a, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("agent_id:", a.AgentID, "task_id:", a.TaskID)
	return nil
}
