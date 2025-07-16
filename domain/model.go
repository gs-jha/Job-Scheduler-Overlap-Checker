package domain

import "time"

type TimeRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type OverlapRequest struct {
	Range1 TimeRange `json:"range1"`
	Range2 TimeRange `json:"range2"`
}

type OverlapResponse struct {
	Overlap bool   `json:"overlap"`
	Message string `json:"message"`
}
