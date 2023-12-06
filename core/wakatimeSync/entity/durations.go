package entity

import "time"

type DurationsRequest struct {
	Date       string `json:"date,omitempty"`
	Project    string `json:"project,omitempty"`
	Branches   string `json:"branches,omitempty"`
	Timeout    string `json:"Timeout,omitempty"`
	WritesOnly bool   `json:"writes_only,omitempty"`
	Timezone   string `json:"timezone,omitempty"`
	SliceBy    string `json:"slice_by,omitempty"`
}
type DurationsResponse struct {
	Branches []string `json:"branches"`
	Data     []struct {
		Color    interface{} `json:"color"`
		Duration float64     `json:"duration"`
		Project  string      `json:"project"`
		Time     float64     `json:"time"`
	} `json:"data"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Timezone string    `json:"timezone"`
}
