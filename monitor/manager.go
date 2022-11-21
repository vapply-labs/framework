package monitor

import (
	"github.com/vapply-labs/framework/jobs"
	"github.com/vapply-labs/framework/tasks"
)

// Allows for flexibility to create different versions of monitors managers if needed
type MonitorsManager interface {
	// Start a monitor task for each given company.
	StartMonitorTasks(companies []*jobs.SupportedCompany) error

	// Stop monitor tasks corresponding to each given company
	StopMonitorTasks(companies []*jobs.SupportedCompany) error

	// Stop all known running tasks.
	StopAll() error
}

type BaseMonitorsManager struct {
	currMonitoredCompanies map[jobs.SupportedCompany][]*tasks.JobMonitorTask
}
