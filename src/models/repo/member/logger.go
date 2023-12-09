package member

import (
	"service-api/src/models/ent"
	"service-api/src/models/repo"
)

type LoggerRepo struct{}

func (LoggerRepo) Write(log ent.MemberAuthorizeLog) (*ent.MemberAuthorizeLog, error) {
	return repo.Query().
		MemberAuthorizeLog.
		Create().
		SetMemberID(log.MemberID).
		SetChannel(log.Channel).
		SetDevice(log.Device).
		SetDeviceDetail(log.DeviceDetail).
		SetRemoteIP(log.RemoteIP).
		SetClientIP(log.ClientIP).
		Save(repo.Ctx())
}
