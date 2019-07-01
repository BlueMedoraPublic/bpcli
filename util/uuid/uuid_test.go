package uuid

import (
    "testing"
)

func TestIsUUID(t *testing.T) {
    if IsUUID("0a65db1a-7171-4318-8e30-548a9ba6b7af") != true {
        t.Errorf("Got false for valid uuid, expected true.")
    }

    if IsUUID("abc") != false {
        t.Errorf("Got true for invalid uuid, expected false.")
    }
}
