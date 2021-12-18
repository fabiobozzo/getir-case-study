package model

import "time"

type Record struct {
	Id         string
	Key        string
	Value      string
	CreatedAt  time.Time
	TotalCount int
}
