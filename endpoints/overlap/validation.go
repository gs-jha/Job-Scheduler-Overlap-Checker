package overlap

import (
	"errors"

	"github.com/gs-jha/Job-Scheduler-Overlap-Checker/domain"
)

func ValidateOverlapRequest(req *domain.OverlapRequest) error {
	if req.Range1.Start.IsZero() || req.Range1.End.IsZero() || req.Range2.Start.IsZero() || req.Range2.End.IsZero() {
		return errors.New("start and end times must be provided")
	}
	if req.Range1.Start.After(req.Range1.End) || req.Range2.Start.After(req.Range2.End) {
		return errors.New("start time must be before end time")
	}
	return nil
}
