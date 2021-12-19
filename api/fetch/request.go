package fetch

import (
	"getir-case-study/pkg/utils"
	"time"
)

type request struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int    `json:"minCount"`
	MaxCount  int    `json:"maxCount"`
}

func (r *request) Validate() (bool, time.Time, time.Time) {
	valid := true

	startTime, err := time.Parse(utils.DateFormat, r.StartDate)
	endTime, err := time.Parse(utils.DateFormat, r.EndDate)

	if err != nil || r.MaxCount < r.MinCount {
		valid = false
	}

	return valid, startTime, endTime
}
