package sdk

import (
	"encoding/json"
	"fmt"
)

// Credential describes a source credential configuration
type Credential struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	URL              string      `json:"url"`
	CredentialTypeID string      `json:"credential_type_id"`
	Parameters       interface{} `json:"parameters"`
	Sources []CredentialSource `json:"sources"`
}

// CredentialSource describes the soure object returned
// inside a Credential object's list of sources
type CredentialSource struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	Stopped    bool   `json:"stopped"`
	Status     string `json:"status"`
	SourceType struct {
		ID     string `json:"id"`
		URL    string `json:"url"`
		DocURL string `json:"doc_url"`
	}
}

// CredentialCreateResponse type describes the json body returned
// by the credential create endpoint
type CredentialCreateResponse struct {
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	URL              string      `json:"url"`
	CredentialTypeID string      `json:"credential_type_id"`
	Parameters       interface{} `json:"parameters"`
}

// CredentialType describes a credential type object
type CredentialType struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Parameters []interface{} `json:"parameters"`
	Credentials []Credential `json:"credentials"`
}

// CredentialTypeTemplate describes a credential type template object
type CredentialTypeTemplate struct {
	Name             string      `json:"name"`
	CredentialTypeID string      `json:"credential_type_id"`
	Parameters       interface{} `json:"parameters"`
}

// GetCredential will return a credential type
func (bp BindPlane) GetCredential(id string) (Credential, error) {
	var c Credential
	body, err := bp.APICall("GET", bp.paths.credentials+"/"+id, nil)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// GetCredentials will return a credential type
func (bp BindPlane) GetCredentials() ([]Credential, error) {
	var c []Credential
	body, err := bp.APICall("get", bp.paths.credentials, nil)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// GetCredentialTemplate will return a Credential configuration
func (bp BindPlane) GetCredentialTemplate(id string) (CredentialTypeTemplate, error) {
	var c CredentialTypeTemplate
	body, err := bp.APICall("get", bp.paths.credentialTypes+"/"+id+"/template", nil)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// GetCredentialType returns a credentials type
func (bp BindPlane) GetCredentialType(id string) (CredentialType, error) {
	var c CredentialType
	body, err := bp.APICall("get", bp.paths.credentialTypes+"/"+id, nil)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// DeleteCredential will delete a configured credential
func (bp BindPlane) DeleteCredential(id string) error {
	_, err := bp.APICall("delete", bp.paths.credentials+"/"+id, nil)
	return err
}

// CreateCredential will configure a credential
func (bp BindPlane) CreateCredential(payload []byte) (CredentialCreateResponse, error) {
	var c CredentialCreateResponse
	body, err := bp.APICall("post", bp.paths.credentials, payload)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// Print will print a Credential object
func (c Credential) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(c, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil

	}

	fmt.Println("id:", c.ID, "credential_type_id", c.CredentialTypeID, "name:", c.Name)
	return nil
}

// Print will print a CredentialCreateResponse object
func (c CredentialCreateResponse) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(c, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("credential_id:", c.ID, "credential_name:", c.Name)
	return nil
}

// Print will print the CredentialType object
func (c CredentialType) Print() error {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf(string(b))
	return nil
}

// Print will print the CredentialTypeTemplate object
func (c CredentialTypeTemplate) Print() error {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf(string(b))
	return nil
}
