package service

import (
	"context"

	"test_task/internal/dto"
	"test_task/internal/entity"
)

type RaribleApiClient interface {
	GetOwnershipByID(ctx context.Context, id string) (*entity.Ownership, error)
	GetNftTraitsRarity(ctx context.Context, collectionId string, properties []dto.PropertyKV) (*dto.TraitRarityResponse, error)
}
