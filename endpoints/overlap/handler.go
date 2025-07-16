package overlap

import (
	"encoding/json"
	"net/http"

	"github.com/gs-jha/Job-Scheduler-Overlap-Checker/domain"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CheckOverlap(w http.ResponseWriter, r *http.Request) {
	var req domain.OverlapRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := ValidateOverlapRequest(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg := "No overlap detected"
	if IsOverlapped(&req) {
		msg = "Overlap detected"
	}

	response := domain.OverlapResponse{
		Overlap: IsOverlapped(&req),
		Message: msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func IsOverlapped(req *domain.OverlapRequest) bool {
	return req.Range1.Start.Before(req.Range2.End) && req.Range2.Start.Before(req.Range1.End)
}
