package authoize

import (
	"service-api/src/models/ent"
	"service-api/src/models/ent/member"
	"service-api/src/models/repo"
)

// AccountExists 判断账号是否存在
func AccountExists(account string) (bool, error) {
	return repo.Query().Member.Query().Where(
		member.Or(
			member.Account(account),
			member.Email(account),
		),
	).Exist(repo.Ctx())
}

// AccountAndEmailExists 判断账号是否存在
func AccountAndEmailExists(account string, email string) bool {
	return repo.Query().Member.Query().Where(
		member.Account(account),
		member.Email(email),
	).ExistX(repo.Ctx())
}

// FirstMemberByAccount 获取一个账号
func FirstMemberByAccount(account string) (*ent.Member, error) {
	return repo.Query().Member.Query().Where(
		member.Or(
			member.Account(account),
			member.Email(account),
		),
	).First(repo.Ctx())
}
