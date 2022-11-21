package monitor

import (
	"github.com/vapply-labs/framework/jobs"
	"github.com/vapply-labs/framework/modules/yelp"
)

// Creates a monitor task for each given company.
// Acts as the main factory function for creating monitor tasks with only the company name.
//
// TODO: If need to support many many websites, initialize struct dynamically at runtime
// See: https://stackoverflow.com/questions/7850140/how-do-you-create-a-new-instance-of-a-struct-from-its-type-at-run-time-in-go
func CreateMonitorTasks(companies []jobs.SupportedCompany) []JobMonitorTask {
	tasks := []JobMonitorTask{}
	for _, company := range companies {
		if company == "Yelp" {
			tasks = append(tasks, &yelp.YelpTask{})
		}

	}

	return tasks
}
