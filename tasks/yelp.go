package tasks

import "github.com/vapply-labs/framework/jobs"

type YelpTask struct {
}

// Scrapes the website content to parse new jobs.
// Each career website should have their own version of this function.
func (t *YelpTask) Scrape() []*jobs.Job {
	return []*jobs.Job{}
}

// Call this callback for each new job found.
func (t *YelpTask) OnNewJob(job *jobs.Job) error {
	return nil
}

// Starts the monitor
func (t *YelpTask) Start() error {
	return nil
}

// Starts the monitor
func (t *YelpTask) IsStarted() bool {
	return false
}
