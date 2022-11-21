package scrapers

import (
	"fmt"

	"github.com/vapply-labs/framework/jobs"
)

type WebScraper interface {
	Scrape() []*jobs.Job
}

type WebScraperPlugin interface {
	Filter(content interface{}) ([]*jobs.Job, error)
	Name() string
}

type JobInfoScraper struct {
	Plugins []WebScraperPlugin
}

func (s *JobInfoScraper) Scrape() ([]*jobs.Job, error) {
	parsedJobs := []*jobs.Job{}
	for _, plugin := range s.Plugins {
		pluginJobs, err := plugin.Filter("dummy content")
		if err != nil {
			return nil, fmt.Errorf("err filtering with plugin %s: %s", plugin.Name(), err)
		}
		parsedJobs = append(parsedJobs, pluginJobs...)
	}

	// Remove duplicate jobs (may not work because all jobs are pointers)
	return jobs.GetUniqueJobs(parsedJobs), nil
}
