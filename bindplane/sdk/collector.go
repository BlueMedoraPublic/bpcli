package sdk

import (
    "fmt"
    "encoding/json"
)

// Collector type describes a collector configuration
type Collector struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Version    string `json:"version"`
	Hostname   string `json:"hostname"`
	Status     string `json:"status"`
	NumSources int    `json:"num_sources"`
}

// CollectorGroup type describes a collector group configuration
type CollectorGroup struct {
    ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Group      bool   `json:"group"`
	Status     string `json:"status"`
	NumSources int    `json:"num_sources"`
	Collectors []Collector `json:"collectors"`
}

// GetCollector will return a collector object
func (bp BindPlane) GetCollector(id string) (Collector, error) {
    var c Collector
    body, err := bp.APICall("get", bp.paths.collectors + "/" + id, nil)
    if err != nil {
        return c, err
    }

    err = json.Unmarshal(body, &c)
    return c, err
}

// GetCollectorGroup will return a collector group object
func (bp BindPlane) GetCollectorGroup(id string) (CollectorGroup, error) {
    var g CollectorGroup
    body, err := bp.APICall("get", bp.paths.collectors + "/" + id, nil)
    if err != nil {
        return g, err
    }

    err = json.Unmarshal(body, &g)
    return g, err
}

// GetCollectors will return an array of collector group objects
func (bp BindPlane) GetCollectors() ([]CollectorGroup, error) {
    var c []CollectorGroup
    body, err := bp.APICall("get", bp.paths.collectors, nil)
    if err != nil {
        return c, err
    }

    err = json.Unmarshal(body, &c)
    return c, err
}

// DeleteCollector will delete a configured collector
func (bp BindPlane) DeleteCollector(id string) error {
    _, err := bp.APICall("delete", bp.paths.collectors + "/" + id, nil)
    return err
}

// Print prints a collector group
func (g CollectorGroup) Print(j bool) error {
    if j == true {
        b, err := json.MarshalIndent(g, "", "  ")
        if err != nil {
            return err
        }

        fmt.Println(string(b))
        return nil
    }

    fmt.Println("id:", g.ID, "status:", g.Status, "type:", "group",
        "name:", g.Name, "num_sources:", g.NumSources)
    return nil
}

// Print prints a collector
func (c Collector) Print(j bool) error {
    if j == true  {
        b, err := json.MarshalIndent(c, "", "  ")
        if err != nil {
            return err
        }

        fmt.Println(string(b))
        return nil
    }

    fmt.Println("id:", c.ID, "status:", c.Status, "type:", "collector",
        "version:", c.Version, "name:", c.Name)
    return nil
}
