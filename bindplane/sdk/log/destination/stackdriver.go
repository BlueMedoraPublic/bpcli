package destination

import (
    "fmt"
    "encoding/json"

    "github.com/pkg/errors"
)

// Stackdriver type represents a Stackdriver destination config
// for bindPlane
type Stackdriver struct {
    Name          string `json:"name"`
	ID            string `json:"id,omitempty"`
    Destination struct {
		ID      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"destination,omitempty"`
    Configuration struct {
        /*Credentials struct {
            Type                    string `json:"type,omitempty"`
			ProjectID               string `json:"project_id,omitempty"`
			PrivateKeyID            string `json:"private_key_id,omitempty"`
			PrivateKey              string `json:"private_key,omitempty"`
			ClientEmail             string `json:"client_email,omitempty"`
			ClientID                string `json:"client_id,omitempty"`
			AuthURI                 string `json:"auth_uri,omitempty"`
			TokenURI                string `json:"token_uri,omitempty"`
			AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url,omitempty"`
			ClientX509CertURL       string `json:"client_x509_cert_url,omitempty"`
        } `json:"credentials,omitempty"`*/
        Credentials map[string]string `json:"credentials,omitempty"`
        Location string `json:"location,omitempty"`
    } `json:"configuration"`
    DestinationTypeID  string `json:"destination_type_id,omitempty"`
    DestinationVersion string `json:"destination_version,omitempty"`
}

// Type returns the destination type id
func (s Stackdriver) Type() string {
    return "stackdriver"
}

// JSON returns the json representation of a Stackdriver type
func (s Stackdriver) JSON() ([]byte, error) {
    return json.Marshal(s)
}

// JSONPretty returns the json representation of a Stackdriver type
func (s Stackdriver) JSONPretty() ([]byte, error) {
    return json.MarshalIndent(s, " ", " ")
}

// Validate validates a Stackdriver configuration
func (s Stackdriver) Validate() error {
    var (
        fail = false
        err = errors.New("Stackdriver config failed validation")
    )

    if s.Name == "" {
        e := err
        err = errors.Wrap(e, "Name is not set")
        fail = true
    }

    /*if s.Configuration.Credentials.Type == "" {
        e := err
        err = errors.Wrap(e, "Type is not set")
        fail = true
    }

    if s.Configuration.Credentials.ProjectID == "" {
        e := err
        err = errors.Wrap(e, "ProjectID is not set")
        fail = true
    }

    if s.Configuration.Credentials.PrivateKeyID == "" {
        e := err
        err = errors.Wrap(e, "PrivateKeyID is not set")
        fail = true
    }

    if s.Configuration.Credentials.PrivateKey == "" {
        e := err
        err = errors.Wrap(e, "PrivateKey is not set")
        fail = true
    }

    if s.Configuration.Credentials.ClientEmail == "" {
        e := err
        err = errors.Wrap(e, "ClientEmail is not set")
        fail = true
    }

    if s.Configuration.Credentials.ClientID == "" {
        e := err
        err = errors.Wrap(e, "ClientID is not set")
        fail = true
    }

    if s.Configuration.Credentials.AuthURI == "" {
        e := err
        err = errors.Wrap(e, "AuthURI is not set")
        fail = true
    }

    if s.Configuration.Credentials.TokenURI == "" {
        e := err
        err = errors.Wrap(e, "TokenURI is not set")
        fail = true
    }

    if s.Configuration.Credentials.AuthProviderX509CertURL == "" {
        e := err
        err = errors.Wrap(e, "AuthProviderX509CertURL is not set")
        fail = true
    }

    if s.Configuration.Credentials.ClientX509CertURL == "" {
        e := err
        err = errors.Wrap(e, "ClientX509CertURL is not set")
        fail = true
    }*/

    if s.Configuration.Location == "" {
        e := err
        err = errors.Wrap(e, "Configuration.Location is not set")
        fail = true
    }

    if s.DestinationTypeID == "" {
        e := err
        err = errors.Wrap(e, "DestinationTypeID is not set")
        fail = true
    }

    if s.DestinationVersion == "" {
        e := err
        err = errors.Wrap(e, "DestinationTypeID is not set")
        fail = true
    }

    // do not check CustomTemplate, not required

    if fail {
        return err
    }
    return nil
}

// Print print a LogDestConfig type
func (s Stackdriver) Print() error {
    b, err := s.JSONPretty()
    if err != nil {
        return err
    }
	fmt.Printf(string(b))
	return nil
}
