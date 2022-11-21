package monitor

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/vapply-labs/framework/jobs"
	"github.com/vapply-labs/framework/tasks"
	"go.uber.org/zap"
)

// Allows for flexibility to create different versions of monitors managers if needed
type MonitorsManager interface {
	// Start a monitor task for each given company.
	// The logger can be optional (specify nil)
	StartMonitorTasks(companies []jobs.SupportedCompany, logger *zap.SugaredLogger) error

	// Stop monitor tasks corresponding to each given company
	// The logger can be optional (specify nil)
	StopMonitorTasks(companies []jobs.SupportedCompany, logger *zap.SugaredLogger) error

	// Stop all known running tasks.
	// The logger can be optional (specify nil)
	StopAll(logger *zap.SugaredLogger) error
}

type BaseMonitorsManager struct {
	currMonitoredCompanies map[jobs.SupportedCompany]tasks.JobMonitorTask
}

func (m *BaseMonitorsManager) StartMonitorTasks(companies []jobs.SupportedCompany, logger *zap.SugaredLogger) error {
	if len(companies) == 0 {
		if logger != nil {
			logger.Debugw("no tasks created", "fn", "StartMonitorTasks")
		}

		return nil
	}

	// Create monitor tasks for only tasks that have not been cached.
	// If a task is cached and already running, ignore it in this function.
	// For others (i.e. cached and not started, create new tasks and start those tasks)
	filteredCompanies := lo.Filter(companies, func(company jobs.SupportedCompany, index int) bool {
		if cachedTask, ok := m.currMonitoredCompanies[company]; ok && cachedTask.IsStarted() {
			// Ignore this company
			return false
		}

		return true
	})

	monitorTasks := tasks.CreateMonitorTasks(filteredCompanies)

	if logger != nil {
		createdTasksLog := fmt.Sprintf("StartMonitorTasks: created %d tasks; %d companies filtered out", len(monitorTasks), len(companies)-len(filteredCompanies))
		logger.Debugw(createdTasksLog, "fn", "StartMonitorTasks")
	}

	for i, task := range monitorTasks {
		err := task.Start()
		if err != nil {
			return fmt.Errorf("initializing monitor task for %s failed with err: %s", companies[i], err)
		}
		// Cache the task
		m.currMonitoredCompanies[companies[i]] = task
	}

	return nil
}
