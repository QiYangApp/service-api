package user

import (
	"context"
	"ent/models"
	accounts "ent/models/accounts"
	"service-api/internal/app/repo"
)

func GetSingleAccountByName(ctx context.Context, account string) (*models.Accounts, error) {
	return repo.Client().Accounts.Query().Where(accounts.AccountEQ(account)).First(ctx)
}
