package pprint

import (
    "fmt"
    "encoding/json"
)

// PrintJSONStringMap takes a map[string]string and
// pretty prints json
func PrintJSONStringMap(m map[string]string) error {
    b, err := json.MarshalIndent(m, " ", "")
    if err != nil {
        return err
    }
    fmt.Println(string(b))
    return nil
}
