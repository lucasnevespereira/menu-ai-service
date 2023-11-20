package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStore struct {
	client *mongo.Client
	conf   UserStoreConfig
}

type UserStoreConfig struct {
	Database   string
	URL        string
	Collection string
}

func NewUserStore(ctx context.Context, conf UserStoreConfig) (*UserStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.URL))
	if err != nil {
		return nil, err
	}

	return &UserStore{client: client, conf: conf}, nil
}
