package monitor

import "github.com/vapply-labs/framework/jobs"

type MonitorsManager interface {
	// Start a monitor task for each given company.
	StartMonitorTasks(companies []*jobs.Job) error

	// Stop monitor tasks corresponding to each given company
	StopMonitorTasks(companies []*jobs.Job) error

	// Stop all known running tasks.
	StopAll() error
}
