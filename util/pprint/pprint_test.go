package pprint

import (
    "testing"
)

func TestPrintJSONStringMapPass(t *testing.T) {
    m := make(map[string]string)
    m["a"] = "a"
    m["b"] = "b"

    if err := PrintJSONStringMap(m); err != nil {
        t.Errorf("expected PrintJSONStringMap() to return a nil error when given a map[string]string")
    }
}

func TestPrintJSONStringMapPass2(t *testing.T) {
    m := map[string]string{}

    if err := PrintJSONStringMap(m); err != nil {
        t.Errorf("expected PrintJSONStringMap() to return a nil error when given an uninitialized map[string]string")
    }
}
