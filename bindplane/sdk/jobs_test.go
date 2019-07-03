package sdk

import (
	"testing"
)

func TestparseStatus(t *testing.T) {
	var job Job

	job.Status = "Complete"
	if job.parseStatus() != 0 {
		t.Errorf("Expected GetStatus() to return 0 when status is 'Complete'")
	}

	job.Status = "Failed"
	if job.parseStatus() != 1 {
		t.Errorf("Expected GetStatus() to return 1 when status is 'Failed'")
	}

	job.Status = "In Progress"
	if job.parseStatus() != 2 {
		t.Errorf("Expected GetStatus() to return 2 when status is 'In Progress'")
	}

	job.Status = "Testing Connection to Source"
	if job.parseStatus() != 2 {
		t.Errorf("Expected GetStatus() to return 1 when status is 'Testing Connection to Source'")
	}

	job.Status = "Queued for completion"
	if job.parseStatus() != 3 {
		t.Errorf("Expected GetStatus() to return 2 when status is 'Queued for completion'")
	}

	job.Status = "some fake message"
	if job.parseStatus() != 100 {
		t.Errorf("Expected GetStatus() to return 100 when status is not defined in GetStatus()")
	}
}
