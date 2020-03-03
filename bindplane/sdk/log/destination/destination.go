package destination

import (
    "strings"
    "encoding/json"

    "github.com/pkg/errors"
)

// Destination interface represents a destination
// configuraton for BindPlane
type Destination interface{
    Validate() error
    JSON() ([]byte, error)
    JSONPretty() ([]byte, error)
    Type() string
    Print() error
}

// New returns a new Destination configurartion for a
// given destination type id
func New(destType string, c []byte) (Destination, error) {
    destType = strings.ToLower(destType)

    if destType == "stackdriver" {
        sd := Stackdriver{}
        err := json.Unmarshal(c, &sd)
        return sd, errors.Wrap(err, string(c))
    }

    if destType == "newrelic" {
        nr := NewRelic{}
        err := json.Unmarshal(c, &nr)
        return nr, err
    }

    err := errors.New(destType + " is an invalid destination type, or not supported by this package")
    return nil, err
}
