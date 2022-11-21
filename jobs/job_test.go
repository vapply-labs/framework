package jobs_test

import (
	"testing"

	"github.com/vapply-labs/framework/jobs"
)

func TestGetUniqueJobs(t *testing.T) {
	hasDuplicates := []*jobs.Job{
		{
			Company:  "Yelp",
			Title:    "Software Engineering Intern",
			Season:   "N/A",
			Location: "N/A",
			IsRemote: true,
			Link:     "N/A",
		},
		{
			Company:  "Yelp",
			Title:    "Software Engineering Intern",
			Season:   "N/A",
			Location: "N/A",
			IsRemote: true,
			Link:     "N/A",
		},
	}

	filtered := jobs.GetUniqueJobs(hasDuplicates)
	if len(filtered) != 1 {
		t.Fatalf("GetUniqueJobs failed to filter duplicates in a jobs array")
	}
}
