// +build integration

package sdk

import (
    "strings"
    "testing"
)

var testSource = strings.TrimRight(`{
  "name": "microsoft_azure_virtualmachines",
  "source_type": "microsoft_azure_virtualmachines",
  "collector_id": "fake",
  "collection_interval": 60,
  "credentials": {
    "credentials": "fake"
  },
  "configuration": {
    "collection_time_grain": "1",
    "filter_by_resource_group_type": "whitelist",
    "filter_by_tags_group_type": "whitelist",
    "http_request_timeout": 30,
    "maximum_http_retry_time": 45,
    "monitor_metric_collection_level": "kpi",
    "port": 443,
    "proxy_password": "",
    "proxy_port": "",
    "proxy_username": "",
    "ssl_config": "Verify"
  }
}`, "\r\n")

func TestListGetSources(t *testing.T) {
	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	sources, err := bp.GetSources()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	for _, source := range sources {
		if len(source.ID) == 0 {
			t.Errorf("Expected source id length to not be 0")
		}

		// test get source function
		_, err := bp.GetSource(source.ID)
		if err != nil {
			t.Errorf("Expected GetSource to return source with id " + source.ID)
		}
	}
}

func TestGetSourceTemplate(t *testing.T) {
	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	_, err = bp.GetSourceTemplate("microsoft_azure_virtualmachines")
	if err != nil {
		t.Errorf("Got error while requesting a source template\n" + err.Error())
	}
}

func TestGetListSourceType(t *testing.T) {
	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	_, err = bp.GetSourceType("microsoft_azure_virtualmachines")
	if err != nil {
		t.Errorf("Got error while requesting a source type\n" + err.Error())
	}

	_, err = bp.ListSourceTypes()
	if err != nil {
		t.Errorf("Got error while listing sources\n" + err.Error())
	}
}

func TestCreateSource(t *testing.T) {
	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	_, err = bp.CreateSource([]byte(testSource))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestDeleteSource(t *testing.T) {
	// non destructive delete
	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	_, err = bp.DeleteSource("fake source id")
	if err == nil {
		t.Errorf(err.Error())
		return
	}
}
