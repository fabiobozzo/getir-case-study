package fetch

import (
	"bytes"
	"context"
	"getir-case-study/pkg/db"
	"getir-case-study/pkg/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

type MockReader struct{}

func (r *MockReader) RecordsByDateAndCountRange(
	ctx context.Context,
	dateRange db.DateRange,
	countRange db.CountRange,
) ([]model.Record, error) {
	return []model.Record{
		{
			Id:         "id-1",
			Key:        "key-1",
			Value:      "value-1",
			CreatedAt:  time.Date(2016, 01, 29, 0, 0, 0, 0, time.UTC),
			TotalCount: 100,
		},
		{
			Id:         "id-2",
			Key:        "key-2",
			Value:      "value-2",
			CreatedAt:  time.Date(2016, 02, 1, 0, 0, 0, 0, time.UTC),
			TotalCount: 420,
		},
	}, nil
}

func TestFetchSuccess(t *testing.T) {
	testee := NewHandler(logrus.New(), &MockReader{})

	reader := bytes.NewReader([]byte(`{"startDate":"2016-01-26", "endDate":"2016-02-02", "minCount": 2700, "maxCount": 3000}`))

	req, err := http.NewRequest("POST", "/fetch", reader)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testee.Handle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"code":0,"msg":"success","records":[{"key":"key-1","createdAt":"2016-01-29T00:00:00.000Z","totalCount":100},{"key":"key-2","createdAt":"2016-02-01T00:00:00.000Z","totalCount":420}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestFetchInvalid(t *testing.T) {
	testee := NewHandler(logrus.New(), &MockReader{})

	reader := bytes.NewReader([]byte(`{"startDate":"2016-01-26", "endDate":"2015-02-02", "minCount": 30000, "maxCount": 3000}`))

	req, err := http.NewRequest("POST", "/fetch", reader)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testee.Handle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"code":1,"msg":"the input provided is invalid"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
