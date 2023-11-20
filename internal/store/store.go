package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Client *mongo.Client
	Conf   Config
}

type Config struct {
	URL      string
	UsersCol string
	MenusCol string
}

func NewStoreClient(ctx context.Context, conf Config) (*Store, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.URL))
	if err != nil {
		return nil, err
	}

	return &Store{Client: client, Conf: conf}, nil
}
