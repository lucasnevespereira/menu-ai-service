package services

import (
	"menu-ai-service/internal/store"
)

type MenuService interface{}

type MenuServiceImpl struct {
	store *store.Store
}

// Enforces implementation of interface at compile time
var _ MenuService = (*MenuServiceImpl)(nil)

func NewMenuService(store *store.Store) *MenuServiceImpl {
	return &MenuServiceImpl{
		store: store,
	}
}
