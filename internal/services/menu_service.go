package services

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"menu-ai-service/internal/models"
	"menu-ai-service/internal/store"
)

type MenuService interface {
	Create(ctx context.Context, request models.MenuRequest) (*models.Menu, error)
}

type MenuServiceImpl struct {
	store *store.MenuStore
}

// Enforces implementation of interface at compile time
var _ MenuService = (*MenuServiceImpl)(nil)

func NewMenuService(store *store.MenuStore) *MenuServiceImpl {
	return &MenuServiceImpl{
		store: store,
	}
}

func (s *MenuServiceImpl) Create(ctx context.Context, request models.MenuRequest) (*models.Menu, error) {

	row := store.MenuRow{
		ID:           primitive.NewObjectID(),
		Content:      bson.Raw(request.Content),
		ShoppingList: bson.Raw(request.ShoppingList),
		Specs: store.MenuSpecsRow{
			MaxCalories: request.Specs.MaxCalories,
			MaxCarbs:    request.Specs.MaxCarbs,
			MaxProteins: request.Specs.MaxProteins,
			MaxFats:     request.Specs.MaxFats,
			Allergies:   request.Specs.Allergies,
		},
	}

	inserted, err := s.store.Create(ctx, row)
	if err != nil {
		return nil, err
	}

	return &models.Menu{
		ID:           inserted.ID.String(),
		Content:      json.RawMessage(inserted.Content),
		ShoppingList: json.RawMessage(inserted.ShoppingList),
		Specs: models.MenuSpecs{
			MaxCalories: inserted.Specs.MaxCalories,
			MaxCarbs:    inserted.Specs.MaxCarbs,
			MaxProteins: inserted.Specs.MaxProteins,
			MaxFats:     inserted.Specs.MaxFats,
			Allergies:   inserted.Specs.Allergies,
		},
	}, nil
}
