package overlap_test

// write tests for the overlap package handler file
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gs-jha/Job-Scheduler-Overlap-Checker/domain"
	"github.com/gs-jha/Job-Scheduler-Overlap-Checker/endpoints/overlap"
)

func TestCheckOverlap(t *testing.T) {
	handler := overlap.NewHandler()

	tests := []struct {
		name       string
		request    domain.OverlapRequest
		wantStatus int
		wantBody   string
	}{
		{
			name: "No Overlap",
			request: domain.OverlapRequest{
				Range1: domain.TimeRange{Start: time.Date(2025, 10, 1, 10, 0, 0, 0, time.UTC), End: time.Date(2025, 10, 1, 12, 0, 0, 0, time.UTC)},
				Range2: domain.TimeRange{Start: time.Date(2025, 10, 1, 13, 0, 0, 0, time.UTC), End: time.Date(2025, 10, 1, 15, 0, 0, 0, time.UTC)},
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"overlap":false,"message":"No overlap detected"}`,
		},
		{
			name: "Overlap Detected",
			request: domain.OverlapRequest{
				Range1: domain.TimeRange{Start: time.Date(2025, 10, 1, 10, 0, 0, 0, time.UTC), End: time.Date(2025, 10, 1, 12, 0, 0, 0, time.UTC)},
				Range2: domain.TimeRange{Start: time.Date(2025, 10, 1, 11, 0, 0, 0, time.UTC), End: time.Date(2025, 10, 1, 13, 0, 0, 0, time.UTC)},
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"overlap":true,"message":"Overlap detected"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.request)
			req, err := http.NewRequest("POST", "/api/v1/check-overlap", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("could not create request %v", err)
			}
			req.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			handler.CheckOverlap(recorder, req)

			if recorder.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, recorder.Code)
			}
		})
	}
}
