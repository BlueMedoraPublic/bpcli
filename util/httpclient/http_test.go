package httpclient

import (
	"testing"
)

func TestStatusIs20X(t *testing.T) {
	// should return false
	if StatusIs20X(500) == true {
		t.Errorf("Expected StatusValid(500) to return false, got true")
	}

	if StatusIs20X(400) == true {
		t.Errorf("Expected StatusValid(500) to return false, got true")
	}

	if StatusIs20X(300) == true {
		t.Errorf("Expected StatusValid(500) to return false, got true")
	}

	// should return true
	if StatusIs20X(201) == false {
		t.Errorf("Expected StatusValid(201) to return true, got false")
	}

	if StatusIs20X(200) == false {
		t.Errorf("Expected StatusValid(200) to return true, got false")
	}
}

func TestCreateRequest(t *testing.T) {
	req, err := CreateRequest("POST", "https://test.com", []byte("payload"), "token")
	if err != nil {
		t.Errorf("Expected CreateRequest() to NOT return an error, got " + err.Error())
		return
	}

	if req.Header.Get("x-bindplane-api-key") != "token" {
		t.Errorf("Expected CreateRequest() to return a http request with header x-bindplane-api-key='token'")
	}

	if req.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Expected CreateRequest() to return a http request with header Content-Type='application/json'")
	}

}
