package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"menu-ai-service/internal/models"
	"menu-ai-service/internal/store"
)

type MenuService interface {
	Save(ctx context.Context, request models.MenuSaveRequest) (*models.Menu, error)
	DeleteByID(ctx context.Context, menuID string) error
	GetByUserID(ctx context.Context, userID string) ([]*models.Menu, error)
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

func (s *MenuServiceImpl) Save(ctx context.Context, request models.MenuSaveRequest) (*models.Menu, error) {
	row := &store.MenuRow{
		ID:           primitive.NewObjectID(),
		Content:      request.Content,
		ShoppingList: request.ShoppingList,
		Specs: &store.MenuSpecsRow{
			MaxCalories: request.Specs.MaxCalories,
			MaxCarbs:    request.Specs.MaxCarbs,
			MaxProteins: request.Specs.MaxProteins,
			MaxFats:     request.Specs.MaxFats,
			Allergies:   request.Specs.Allergies,
		},
		UserID: request.UserID,
	}

	inserted, err := s.store.Insert(ctx, row)
	if err != nil {
		return nil, err
	}

	return &models.Menu{
		ID:           inserted.ID.Hex(),
		Content:      inserted.Content,
		ShoppingList: inserted.ShoppingList,
		Specs: models.MenuSpecs{
			MaxCalories: inserted.Specs.MaxCalories,
			MaxCarbs:    inserted.Specs.MaxCarbs,
			MaxProteins: inserted.Specs.MaxProteins,
			MaxFats:     inserted.Specs.MaxFats,
			Allergies:   inserted.Specs.Allergies,
		},
		UserID: inserted.UserID,
	}, nil
}

func (s *MenuServiceImpl) DeleteByID(ctx context.Context, menuID string) error {
	err := s.store.Delete(ctx, menuID)
	if err != nil {
		return err
	}
	return nil
}

func (s *MenuServiceImpl) GetByUserID(ctx context.Context, userID string) ([]*models.Menu, error) {
	rowMenus, err := s.store.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return toMenus(rowMenus), nil
}

func toMenus(rowMenus []*store.MenuRow) []*models.Menu {
	var menusList []*models.Menu

	for _, rowMenu := range rowMenus {
		menusList = append(menusList, &models.Menu{
			ID:           rowMenu.ID.Hex(),
			Content:      rowMenu.Content,
			ShoppingList: rowMenu.ShoppingList,
			Specs: models.MenuSpecs{
				MaxCalories: rowMenu.Specs.MaxCalories,
				MaxCarbs:    rowMenu.Specs.MaxCarbs,
				MaxProteins: rowMenu.Specs.MaxProteins,
				MaxFats:     rowMenu.Specs.MaxFats,
				Allergies:   rowMenu.Specs.Allergies,
			},
			UserID: rowMenu.UserID,
		})
	}

	return menusList
}
