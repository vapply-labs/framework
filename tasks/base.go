package tasks

// Monitors a careers website, and when it finds a new job, call "OnNewJob"
type JobMonitorTask interface {
	// Scrapes the website content to parse new jobs.
	Scrape() []*Job

	// Call this callback for each new job found.
	OnNewJob(job *Job) error
}
