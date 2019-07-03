package sdk

import (
	"encoding/json"
	"fmt"
)

// Job describes a job object
type Job struct {
	ID      string      `json:"id"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	URL     string      `json:"url"`
	Result  interface{} `json:"result"`
}

// ListJobs will return an array of collector objects
func (bp BindPlane) ListJobs() ([]Job, error) {
	var c []Job
	body, err := bp.APICall("GET", bp.paths.jobs, nil)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// GetJob will return a job object using an interface{} for
// the result key
func (bp BindPlane) GetJob(id string) (Job, error) {
	var c Job
	body, err := bp.APICall("GET", bp.paths.jobs+"/"+id, nil)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	return c, err
}

// Print will print a Job object
func (job Job) Print(j bool) error {
	if j == true {
		b, err := json.MarshalIndent(job, "", "  ")
		if err != nil {
			return err
		}
		fmt.Printf(string(b))
		return nil
	}

	fmt.Println("id:", job.ID, "status:", job.Status)
	return nil
}

// GetStatus returns an int representing the job status
func (bp BindPlane) GetStatus(job Job) int {
	x, err := bp.GetJob(job.ID)
	if err != nil {
		return -1
	}
	return x.parseStatus()
}

/*
Returns 0 if JobStatus is "Complete"
Returns 1 if JobStatus is "Failed"
Returns 2 if JobStatus is "In Progress" or "Testing Connection to Source"
Returns 3 if JobStatus is "Queued for completion"
Returns 100 if JobStatus is not defined in the function
*/
func (job Job) parseStatus() int {
	if job.Status == "Complete" {
		return 0
	}
	if job.Status == "Failed" {
		return 1
	}
	if job.Status == "In Progress" || job.Status == "Testing Connection to Source" {
		return 2
	}
	if job.Status == "Queued for completion" {
		return 3
	}
	return 100
}
