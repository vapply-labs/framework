package yelp

import (
	"fmt"

	"github.com/vapply-labs/framework/jobs"
	"github.com/vapply-labs/framework/scrapers"
)

// const yelpEngCareers = "https://www.yelp.careers/us/en/c/engineering-jobs?from=10&s=1"

type YelpScraper struct {
	Plugins []scrapers.WebScraperPlugin
}

// type JobInfoScraper struct {
// 	Plugins []WebScraperPlugin
// }

func (s *YelpScraper) Scrape() ([]*jobs.Job, error) {
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
