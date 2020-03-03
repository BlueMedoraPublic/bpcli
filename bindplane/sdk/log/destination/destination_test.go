package destination

import (
    "testing"
)

func TestNew(t *testing.T) {
    // only test bad config, other test functions will test New()
    // due to their nature
    _, err := New("fakedest-abcd", []byte("{}"))
    if err == nil {
        t.Errorf("expected New(fakedest-abcd) to return an error, got nil")
    }
}

func TestValidate(t *testing.T) {
    for _, d := range destinationTypes() {
        config, err := New(d, newConfigBytes(d))
        if err != nil {
            t.Errorf("expected New(" + d + ") to return a nil error, got: " + err.Error())
        } else {
            if err := config.Validate(); err != nil {
                t.Errorf("expected " + d + ".Validate() to return a nil error, got: " + err.Error())
            }
        }
    }
}

func TestJSON(t *testing.T) {
    for _, d := range destinationTypes() {
        config, err := New(d, newConfigBytes(d))
        if err != nil {
            t.Errorf("expected New(" + d + ") to return a nil error, got: " + err.Error())
        }
        if _, err := config.JSON(); err != nil {
            t.Errorf("exected " + d + ".JSON() to return a nil error, got: " + err.Error())
        }
    }
}

func TestJSONPretty(t *testing.T) {
    for _, d := range destinationTypes() {
        config, err := New(d, newConfigBytes(d))
        if err != nil {
            t.Errorf("expected New(" + d + ") to return a nil error, got: " + err.Error())
        }
        if _, err := config.JSONPretty(); err != nil {
            t.Errorf("exected " + d + ".JSON() to return a nil error, got: " + err.Error())
        }
    }
}

func destinationTypes() []string {
    return []string{"stackdriver", "newrelic"}
}

func newConfigBytes(t string) []byte {
    if t == "stackdriver" {
        return []byte(`
{
  "name": "test-json",
  "configuration": {
    "credentials": {
  "type": "service_account",
  "project_id": "bpcli-dev",
  "private_key_id": "redacted",
  "private_key": "redacted",
  "client_email": "redacted@redacted.gserviceaccount.com",
  "client_id": "redacted",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/redacted%40redacted.iam.gserviceaccount.com"
},
    "location": "us-west1"
  },
  "destination_type_id": "stackdriver",
  "destination_version": "1.3.2"
}`)
    }

    if t == "newrelic" {
        return []byte(`
{
"name": "test-json",
"configuration": {
"api_key": "test_api_key"
},
"destination_type_id": "newrelic",
"destination_version": "0.0.1",
"custom_template": ""
}`)
    }

    return nil
}
