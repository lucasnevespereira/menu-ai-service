package services

import (
	"context"
	"log"
	"menu-ai-service/configs"
	"menu-ai-service/internal/store"
)

type Services struct {
	MenuService *MenuServiceImpl
}

func InitServices(config configs.Config) *Services {
	menuStore, err := store.NewMenuStore(context.Background(),
		store.MenuStoreConfig{
			Database:   config.DbName,
			URL:        config.DbUrl,
			Collection: config.DbMenusCollection,
		},
	)
	if err != nil {
		log.Printf("InitServices: could not init menu store: %v \n", err)
	}

	return &Services{
		MenuService: NewMenuService(menuStore),
	}

}
