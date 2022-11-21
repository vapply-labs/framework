package monitor

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/vapply-labs/framework/jobs"
	"go.uber.org/zap"
)

// Allows for flexibility to create different versions of monitors managers if needed
type MonitorsManager interface {
	// Start a monitor task for each given company. If a task does not already exist for that company,
	// create that task.
	StartMonitorTasks(companies []jobs.SupportedCompany) error

	// Stop monitor tasks corresponding to each given company
	// Returns the companies that had a successful stop (i.e. exist and were running)
	StopMonitorTasks(companies []jobs.SupportedCompany) ([]jobs.SupportedCompany, error)

	// Stop all known running tasks.
	// Returns the companies that had a successful stop (i.e. exist and were running)
	StopAll() ([]jobs.SupportedCompany, error)
}

type BaseMonitorsManager struct {
	taskCache map[jobs.SupportedCompany]JobMonitorTask
	logger    *zap.SugaredLogger
}

// Start a monitor task for each given company. If a task does not already exist for that company,
// create that task.
func (m *BaseMonitorsManager) StartMonitorTasks(companies []jobs.SupportedCompany) error {
	if len(companies) == 0 {
		if m.logger != nil {
			m.logger.Debugw("no tasks created", "fn", "StartMonitorTasks")
		}

		return nil
	}

	// Create monitor tasks for only tasks that have not been cached.
	// If a task is cached and already running, ignore it in this function.
	// For others (i.e. cached and not started, create new tasks and start those tasks)
	filteredCompanies := lo.Filter(companies, func(company jobs.SupportedCompany, index int) bool {
		if cachedTask, ok := m.taskCache[company]; ok && cachedTask.IsRunning() {
			// Ignore this company
			return false
		}

		return true
	})

	monitorTasks := CreateMonitorTasks(filteredCompanies)

	if m.logger != nil {
		createdTasksLog := fmt.Sprintf("StartMonitorTasks: created %d tasks; %d companies filtered out", len(monitorTasks), len(companies)-len(filteredCompanies))
		m.logger.Debugw(createdTasksLog, "fn", "StartMonitorTasks")
	}

	for i, task := range monitorTasks {
		err := task.Start()
		if err != nil {
			return fmt.Errorf("initializing monitor task for %s failed with err: %s", companies[i], err)
		}

		// Cache the task
		m.taskCache[companies[i]] = task
	}

	return nil
}

// Ensures that all tasks for the companies are stopped if they exist.
// Throws error when we fail to properly stop a running task.
func (m *BaseMonitorsManager) StopMonitorTasks(companies []jobs.SupportedCompany) ([]jobs.SupportedCompany, error) {
	stopped := []jobs.SupportedCompany{}
	for _, company := range companies {
		if cachedTask, ok := m.taskCache[company]; ok && cachedTask.IsRunning() {
			err := cachedTask.Stop()
			if err != nil {
				return nil, fmt.Errorf("err stopping task for %s: %s", company, err)
			}
			stopped = append(stopped, company)
		}
	}

	return stopped, nil
}

// Ensures that all tasks for the companies are stopped if they exist.
// Throws error when we fail to properly stop a running task.
func (m *BaseMonitorsManager) StopAll() ([]jobs.SupportedCompany, error) {
	stopped := []jobs.SupportedCompany{}
	for company, task := range m.taskCache {
		if task.IsRunning() {
			err := task.Stop()
			if err != nil {
				return nil, fmt.Errorf("err stopping task for %s: %s", company, err)
			}
			stopped = append(stopped, company)
		}
	}

	return stopped, nil
}
