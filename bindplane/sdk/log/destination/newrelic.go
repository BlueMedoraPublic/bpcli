package destination

import (
    "fmt"
    "encoding/json"

    "github.com/pkg/errors"
)

// NewRelic type represents a NewRelic destination config
// for BindPlane
type NewRelic struct {
    Name          string `json:"name"`
    ID          string `json:"id,omitempty"`
    Destination struct {
        ID      string `json:"id,omitempty"`
        Name    string `json:"name,omitempty"`
        Version string `json:"version,omitempty"`
    } `json:"destination,omitempty"`
    Configuration struct {
        APIKey string `json:"api_key,omitempty"`
    } `json:"configuration,omitempty"`
    DestinationTypeID  string `json:"destination_type_id,omitempty"`
    DestinationVersion string `json:"destination_version,omitempty"`
    CustomTemplate     string `json:"custom_template,omitempty"`
}

// Type returns the destination type id
func (n NewRelic) Type() string {
    return "newrelic"
}

// JSON returns the json representation of a NewRelic type
func (n NewRelic) JSON() ([]byte, error) {
    return json.Marshal(n)
}

// JSONPretty returns the json representation of a NewRelic type
func (n NewRelic) JSONPretty() ([]byte, error) {
    return json.MarshalIndent(n, " ", " ")
}

// Validate validates a NewRelic configuration
func (n NewRelic) Validate() error {
    var (
        fail = false
        err = errors.New("NewRelic config failed validation")
    )

    if n.Name == "" {
        err = errors.Wrap(err, "Name is not set")
    }

    if n.Configuration.APIKey == "" {
        err = errors.Wrap(err, "Configuration.APIKey is not set")
    }

    if n.DestinationTypeID == "" {
        err = errors.Wrap(err, "DestinationTypeID is not set")
    }

    if n.DestinationVersion == "" {
        err = errors.Wrap(err, "DestinationVersion is not set")
    }

    // do not check CustomTemplate, not required

    if fail {
        return err
    }
    return nil
}

// Print print a LogDestConfig type
func (n NewRelic) Print() error {
    b, err := n.JSONPretty()
    if err != nil {
        return err
    }
	fmt.Printf(string(b))
	return nil
}
