package user

import (
	"context"
	"service-api/internal/app/repo"
)

func SetUserLanguage(ctx context.Context, uId int64, language string) error {
	u, err := repo.Client().User.Get(ctx, uId)
	if err != nil {
		return err
	}

	u.Language = language
	u, err = repo.Client().User.UpdateOne(u).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
