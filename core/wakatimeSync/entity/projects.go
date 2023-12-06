package entity

import "time"

type ProjectsRequest struct {
	Q string `json:"q"`
}
type ProjectsResponse struct {
	Data []struct {
		Badge                        interface{}   `json:"badge"`
		Clients                      []interface{} `json:"clients"`
		Color                        interface{}   `json:"color"`
		CreatedAt                    time.Time     `json:"created_at"`
		HasPublicUrl                 bool          `json:"has_public_url"`
		HumanReadableLastHeartbeatAt string        `json:"human_readable_last_heartbeat_at"`
		Id                           string        `json:"id"`
		LastHeartbeatAt              time.Time     `json:"last_heartbeat_at"`
		Name                         string        `json:"name"`
		Repository                   interface{}   `json:"repository"`
		Url                          string        `json:"url"`
		UrlencodedName               string        `json:"urlencoded_name"`
	} `json:"data"`
	NextPage   int `json:"next_page"`
	Page       int `json:"page"`
	PrevPage   int `json:"prev_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}
