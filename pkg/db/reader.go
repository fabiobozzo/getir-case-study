package db

import (
	"context"
	"getir-case-study/pkg/model"
	"time"
)

type DateRange struct {
	Start, End time.Time
}

type CountRange struct {
	Min, Max int
}

type Reader interface {
	RecordsByDateAndCountRange(context.Context, DateRange, CountRange) ([]model.Record, error)
}
