package member

import (
	"service-api/src/models/ent"
	"service-api/src/models/repo"
)

type Model struct {
	repo.BaseModel
}

func (m *Model) Create(entity *ent.Member) *ent.Member {
	return m.NewQuery().
		Member.
		Create().
		SetEmail(entity.Email).
		SetMobile(entity.Mobile).
		SetNickname(entity.Nickname).
		SetAvatar(entity.Avatar).
		SetCreateTime(entity.CreateTime).
		SetState(entity.State).
		SetPassword(entity.Password).
		SetPasswordSing(entity.PasswordSing).
		SaveX(m.Ctx())
}
