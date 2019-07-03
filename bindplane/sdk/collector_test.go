package sdk

import (
	"os"
	"testing"
)

func TestGetCollector(t *testing.T) {
	if liveTest() == false {
		return
	}

	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// test bad collector id
	_, err = bp.GetCollector("abc")
	if err == nil {
		t.Errorf("Expected an api error when calling GetCollector() with a bad collector id")
	}
}

func TestGetCollectors(t *testing.T) {
	if liveTest() == false {
		return
	}

	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	c, err := bp.GetCollectors()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	for _, group := range c {
		if len(group.Name) == 0 {
			t.Errorf("Expected collector group name to not be length 0")
		}

		if len(group.ID) == 0 {
			t.Errorf("Expected collector group id to not be length 0")
		}

		if group.Status != "Active" {
			if group.Status != "Error" {
				t.Errorf("Expected collector group status to be Active or Error, got: " + group.Status)
			}
		}

		for _, collector := range group.Collectors {
			if len(collector.Name) == 0 {
				t.Errorf("Expected collector name to not be length 0")
			}

			if len(collector.ID) == 0 {
				t.Errorf("Expected collector id to not be length 0")
			}

			if collector.Status != "Active" {
				if collector.Status != "Error" {
					t.Errorf("Expected collector status to be Active or Error, got: " + group.Status)
				}
			}
		}
	}
}

func TestDeleteCollector(t *testing.T) {
	// non destructive test, expect an error from api
	if liveTest() == false {
		return
	}

	var bp BindPlane
	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = bp.DeleteCollector("fake id here")
	if err == nil {
		t.Errorf("Expected DeleteCollector() to return an error when using a bad id")
	}
}

func liveTest() bool {
	x := os.Getenv("BINDPLANE_LIVE_TEST")
	if x == "1" {
		return true
	}
	return false
}
