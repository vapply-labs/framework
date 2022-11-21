package jobs

import (
	"fmt"

	"github.com/samber/lo"
	"gopkg.in/go-playground/validator.v9"
)

// Master list of supported companies
var supportedCompanies = []string{
	"Yelp",
	"Yahoo",
}

// Seasons for internships
// I.e. SWE Intern for Winter 2023
// Sometimes can be for both, so then should be N/A.
type JobSeason string

const (
	Summer       JobSeason = "Summer"
	Winter       JobSeason = "Winter"
	NotAvailable JobSeason = "N/A"
)

// Checks if a job season is valid
func (s JobSeason) IsValid() bool {
	return s == Summer || s == Winter || s == NotAvailable
}

type SupportedCompany string

// Checks if a company is a supported company (type sensitive)
func (s SupportedCompany) IsValid() bool {
	return lo.Contains(supportedCompanies, string(s))
}

// Universal information to scrape for a new "job"
type Job struct {
	Company  SupportedCompany `validate:"required"`
	Title    string           `validate:"required"`
	Season   JobSeason        `validate:"required"`
	Location string           `validate:"required"` // custom: if not available, "N/A"
	IsRemote bool             `validate:"required"`
	Link     string           `validate:"required"` // custom: if not available, "N/A"
}

// Validates a created job.
func (j *Job) Validate() error {
	validate := validator.New()

	err := validate.Struct(j)
	if err != nil {
		return fmt.Errorf("err validating job: %s", err)
	}

	if !j.Company.IsValid() {
		return fmt.Errorf("not supported company: %s", j.Company)
	}

	if !j.Season.IsValid() {
		return fmt.Errorf("invalid job season: %s", j.Season)
	}

	return nil
}
