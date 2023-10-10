package authoize

import (
	"service-api/src/models/ent/member"
	"service-api/src/models/repo"
)

// ExistsBYPassword 判断用户密码是否正确
func ExistsBYPassword(account string, password string) bool {
	return repo.Query().Member.Query().Where(
		member.Or(
			member.Account(account),
			member.Email(account),
		),
	).Where(member.Password(password)).ExistX(repo.Ctx())
}
