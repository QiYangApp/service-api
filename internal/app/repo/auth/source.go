package auth

import (
	"context"
	"ent/models"
	sourcefield "ent/models/source"
	"service-api/internal/repo"
)

func GetAllSourceByIsActive(ctx context.Context, isActive bool) ([]*models.Source, error) {
	return repo.Client().Source.Query().Where(sourcefield.IsActiveEQ(isActive)).All(ctx)
}
