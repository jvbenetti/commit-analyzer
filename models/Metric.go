package models

type Metric struct {
	Total int            `json:"total"`
	Type  map[string]int `json:"type"`
}
