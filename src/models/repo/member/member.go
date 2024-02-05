package member

import (
	"service-api/src/models/ent"
	"service-api/src/models/repo"
)

// AddSingleByRegister 注册时候添加会员
func AddSingleByRegister(entity *ent.Member) *ent.Member {
	return repo.Query().
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
		SaveX(repo.Ctx())
}
