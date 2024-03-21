package user

import (
	"context"
	"ent/models"
	"ent/models/accounts"
	"service-api/internal/repo"
)

func GetSingleAccountByName(ctx context.Context, username string) (*models.Accounts, error) {
	return repo.Client.Accounts.Query().Where(accounts.AccountEQ(username)).First(ctx)
}
