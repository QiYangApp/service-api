package net

import (
	"encoding/json"
	"wakatimeSync/entity"
)

type ApiInterface interface{}
type Api struct {
	ApiInterface
	Token string
}

func (a *Api) Summary(req entity.SummariesRequest, body entity.SummariesResponse) error {
	return Get(SUMMARY, a.ToMap(req), body)
}

func (a *Api) Heartbeat(req entity.HeartbeatRequest, body entity.HeartbeatResponse) error {
	return Get(HeartBeats, a.ToMap(req), body)
}

func (a *Api) Projects(req entity.ProjectsRequest, body entity.ProjectsResponse) error {
	return Get(PROJECTS, a.ToMap(req), body)
}

func (a *Api) Durations(req entity.DurationsRequest, body entity.DurationsResponse) error {
	return Get(DURATION, a.ToMap(req), body)
}

func (a *Api) UserInfo(req entity.UserRequest, body entity.UserResponse) error {
	return Get(UserInfo, a.ToMap(req), body)
}

func (a Api) ToMap(req interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(req)
	_ = json.Unmarshal(inrec, &inInterface)

	inInterface["token"] = a.Token
	return inInterface
}
