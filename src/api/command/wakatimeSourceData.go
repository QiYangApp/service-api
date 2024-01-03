package command

import (
	"encoding/json"
	"github.com/golang-module/carbon/v2"
	"service-api/src/models/repo"

	"service-api/src/enums"
	"service-api/src/models/repo/wakatimeRepo"
	"wakatimeSync"
	syncEntity "wakatimeSync/entity"
)

type WakatimeSourceData struct {
}

func (*WakatimeSourceData) Run() {

	nodes := wakatimeRepo.AllWakatimeByState(enums.StateEnumOn)

	// 获取当前时间
	now := carbon.Now()

	for _, node := range nodes {
		dataSet := map[string]interface{}{}
		sync := wakatimeSync.New(node.Key)

		summary := &syncEntity.SummariesResponse{}
		_ = sync.Summary(
			syncEntity.SummariesRequest{
				Start: now.StartOfDay().ToDateString(),
				End:   now.AddDays(1).ToDateString(),
			},
			summary,
		)

		dataSet["summary"] = summary

		project := &syncEntity.ProjectsResponse{}
		_ = sync.Projects(
			syncEntity.ProjectsRequest{},
			project,
		)

		dataSet["project"] = project

		heartbeat := &syncEntity.HeartbeatResponse{}
		_ = sync.Heartbeat(
			syncEntity.HeartbeatRequest{
				Date: now.StartOfDay().ToDateString(),
			},
			heartbeat,
		)

		dataSet["heartbeat"] = heartbeat

		durations := &syncEntity.DurationsResponse{}
		_ = sync.Durations(
			syncEntity.DurationsRequest{
				Date: now.StartOfDay().ToDateString(),
			},
			durations,
		)

		dataSet["durations"] = durations

		user := &syncEntity.UserResponse{}
		_ = sync.UserInfo(
			syncEntity.UserRequest{},
			user,
		)

		dataSet["user"] = user

		t, err := json.Marshal(dataSet)
		if err != nil {
			continue
		}

		repo.Query().SourceData.Create().
			SetInfo("111").
			SetType("111").
			SetSubType("1111").
			SetSnapshot(string(t)).
			SetMemberID(node.MemberID).
			SaveX(repo.Ctx())
	}
}
