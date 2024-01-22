package member

import (
	"service-api/src/models/ent"
	"service-api/src/models/ent/member"
	"service-api/src/models/repo"
)

// FindByEmail 获取一个账号
func FindByEmail(email string) (*ent.Member, error) {
	return repo.Query().Member.Query().Where(member.EmailEQ(email)).First(repo.Ctx())
}

// ExistsBYPassword 判断用户密码是否正确
func ExistsBYPassword(account string, password string) bool {
	return repo.Query().Member.Query().Where(member.EmailEQ(account)).Where(member.Password(password)).ExistX(repo.Ctx())
}

// EmailByExists 判断账号是否存在
func EmailByExists(account string) bool {
	return repo.Query().Member.Query().Where(member.EmailEQ(account)).ExistX(repo.Ctx())
}
