package monitor

import (
	"github.com/vapply-labs/framework/jobs"
	"github.com/vapply-labs/framework/tasks"
)

type MonitorsManager interface {
	// Creates a monitor task for each given company.
	CreateMonitorTasks(companies []*jobs.Job) []*tasks.JobMonitorTask

	// Start a monitor task for each given company.
	StartMonitorTasks(companies []*jobs.Job) error

	// Stop monitor tasks corresponding to each given company
	StopMonitorTasks(companies []*jobs.Job) error

	// Stop all known running tasks.
	StopAll() error
}
