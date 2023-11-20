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
	storeClient, err := store.NewStoreClient(context.Background(),
		store.Config{
			URL:      config.DbUrl,
			UsersCol: config.DbUsersCollection,
			MenusCol: config.DbMenusCollection,
		},
	)
	if err != nil {
		log.Printf("could not init store: %v \n", err)
	}

	return &Services{
		MenuService: NewMenuService(storeClient),
	}

}
