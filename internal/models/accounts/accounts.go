package accounts

import (
	"context"
	"service-api/internal/ent"
	"service-api/internal/ent/accounts"
	"service-api/internal/models"
)

func Find(account string, typ Type, ctx context.Context) (*ent.Accounts, error) {
	return models.
		Client().
		Accounts.
		Query().
		Where(accounts.AccountEQ(account)).
		Where(accounts.TypeEQ(typ.UInt())).
		First(ctx)
}

func FindByAccount(account string, ctx context.Context) (*ent.Accounts, error) {
	return models.Client().Accounts.Query().Where(accounts.Account(account)).First(ctx)
}
