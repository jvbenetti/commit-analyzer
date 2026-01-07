package models

type Metric struct {
	Total int            `json:"total"` // Total of commits
	Type  map[string]int `json:"type"`
}
