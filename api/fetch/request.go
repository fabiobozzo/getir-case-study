package fetch

import (
	"getir-case-study/pkg"
	"time"
)

type request struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

func (r *request) ParseStartDate() (time.Time, error) {
	return time.Parse(pkg.DateFormat, r.StartDate)
}

func (r *request) ParseEndDate() (time.Time, error) {
	return time.Parse(pkg.DateFormat, r.EndDate)
}
