package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Request returns a response body and status code
func Request(method string, uri string, payload []byte, token string) ([]byte, error) {
	req, err := CreateRequest(method, uri, payload, token)
	if err != nil {
		return nil, err
	}

	body, status, err := performRequest(req)
	if err != nil {
		return nil, errors.Wrap(err, "error while making request to " + uri)
	}

	if StatusIs20X(status) == false {
		return body, apiErrorHelper(uri, status, body)
	}
	return body, err
}

// CreateRequest returns an http request with headers
func CreateRequest(method string, uri string, payload []byte, token string) (*http.Request, error) {
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-bindplane-api-key", token)
	return req, err
}

// PerformRequest performs an HTTP request and returns a
// response body and status code
func performRequest(req *http.Request) ([]byte, int, error) {
	if err := validMethod(req.Method); err != nil {
		return nil, 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	return body, resp.StatusCode, err
}

// apiErrorHelper formats an error message
func apiErrorHelper(uri string, status int, respBody []byte) error {
	return errors.New(uri + " returned " + strconv.Itoa(status) + "\n" + string(respBody))
}

// StatusIs20X takes a status code, returns true if status
// is a 20X
func StatusIs20X(status int) bool {
	switch status {
	case 200:
		return true
	case 201:
		return true
	case 202:
		return true
	case 203:
		return true
	case 204:
		return true
	default:
		return false
	}
}

func validMethod(method string) error {
	switch method {
	case "GET", "POST", "PUT", "PATCH", "DELETE":
		return nil
	default:
		return errors.New("invalid http method '" + method + "'")
	}
}
