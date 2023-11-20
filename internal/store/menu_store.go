package store

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MenuRow struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Content      bson.Raw           `bson:"content, omitempty"`
	ShoppingList bson.Raw           `bson:"shopping_list, omitempty"`
	Specs        MenuSpecsRow       `son:"specs, omitempty"`
}

type MenuSpecsRow struct {
	MaxCalories string   `bson:"max_calories,omitempty"`
	MaxCarbs    string   `bson:"max_carbs,omitempty"`
	MaxProteins string   `bson:"max_proteins,omitempty"`
	MaxFats     string   `bson:"max_fats,omitempty"`
	Allergies   []string `bson:"allergies,omitempty"`
}

type MenuStore struct {
	client *mongo.Client
	conf   MenuStoreConfig
}

type MenuStoreConfig struct {
	Database   string
	URL        string
	Collection string
}

func NewMenuStore(ctx context.Context, conf MenuStoreConfig) (*MenuStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.URL))
	if err != nil {
		return nil, err
	}

	return &MenuStore{client: client, conf: conf}, nil
}

func (s *MenuStore) Create(ctx context.Context, row MenuRow) (*MenuRow, error) {
	menuCol := s.client.Database(s.conf.Database).Collection(s.conf.Collection)
	one, err := menuCol.InsertOne(ctx, row)

	if err != nil {
		return nil, errors.Wrapf(err, "inserting menu")
	}

	log.Printf("inserted menu with %s \n", one.InsertedID)

	return &row, nil
}
