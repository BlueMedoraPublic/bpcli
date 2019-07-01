package sdk

import (
    "strings"
    "testing"
)

// testCred represent the azure credential type
var testCred = strings.TrimRight(`{
  "name": "ci test",
  "credential_type_id": "515a6b7c-a570-4274-becf-ce651a13e281",
  "parameters": {
    "subscription_id": "abc",
    "tenant_id": "abc",
    "client_id": "abc",
    "client_secret": "abc"
  }
}`, "\r\n")

func TestGetCredential(t *testing.T) {
    if liveTest() == false {
        return
    }

    var bp BindPlane
    err := bp.Init()
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    _, err = bp.GetCredential("fake id here")
    if err == nil {
        t.Errorf("Expected GetCredential() to return an error when using a bad id")
    }
}

func TestGetCredentials(t *testing.T) {
    if liveTest() == false {
        return
    }

    var bp BindPlane
    err := bp.Init()
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    creds, err := bp.GetCredentials()
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    for _, credential := range creds {
        if len(credential.CredentialTypeID) == 0 {
            t.Errorf("Expected credential type id not to be length 0")
        }

        if len(credential.ID) == 0 {
            t.Errorf("Expected credential id not to be length 0")
        }
    }
}

func TestGetCredentialTemplate(t *testing.T) {
    if liveTest() == false {
        return
    }

    var bp BindPlane
    err := bp.Init()
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    c, err := bp.GetCredentialTemplate("515a6b7c-a570-4274-becf-ce651a13e281")
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    if c.CredentialTypeID != "515a6b7c-a570-4274-becf-ce651a13e281" {
        t.Errorf("Expected GetCredentialTemplate(\"515a6b7c-a570-4274-becf-ce651a13e281\") to return 'id 515a6b7c-a570-4274-becf-ce651a13e281'")
    }
}

func TestGetCredentialType(t *testing.T) {
    if liveTest() == false {
        return
    }

    var bp BindPlane
    err := bp.Init()
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    c, err := bp.GetCredentialType("515a6b7c-a570-4274-becf-ce651a13e281")
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    if c.ID != "515a6b7c-a570-4274-becf-ce651a13e281" {
        t.Errorf("Expected GetCredentialType(\"515a6b7c-a570-4274-becf-ce651a13e281\") to return 'id 515a6b7c-a570-4274-becf-ce651a13e281'")
    }
}

func TestCreateDeleteCredential(t *testing.T) {
    if liveTest() == false {
        return
    }

    var bp BindPlane
    err := bp.Init()
    if err != nil {
        t.Errorf(err.Error())
        return
    }

    err = bp.DeleteCredential("some invalid id")
    if err == nil {
        t.Errorf("Expected DeleteCredential() to return an error when passed an invalid id")
    }

    // create a real credential, and then delete it
    resp, err := bp.CreateCredential([]byte(testCred))
    if err != nil {
        t.Errorf("Expected to create a credential, got: " + err.Error())
        t.Errorf(testCred)
        return
    }

    err = bp.DeleteCredential(resp.ID)
    if err != nil {
        t.Errorf("Expected to delete credential, got: " + err.Error())
    }

}
