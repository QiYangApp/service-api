package user

import (
	"context"
	"service-api/internal/ent"
	"service-api/internal/models"
)

func FindById(ctx context.Context, id int) (*ent.User, error) {
	return models.Client().User.Get(ctx, id)
}
