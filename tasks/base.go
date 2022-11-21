package tasks

import "github.com/vapply-labs/framework/jobs"

// Monitors a careers website, and when it finds a new job, call "OnNewJob"
type JobMonitorTask interface {
	// Scrapes the website content to parse new jobs.
	// Each career website should have their own version of this function.
	Scrape() []*jobs.Job

	// Call this callback for each new job found.
	OnNewJob(job *jobs.Job) error

	// Starts the monitor
	Start() error

	IsStarted() bool
}
