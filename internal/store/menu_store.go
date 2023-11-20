package store

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MenuRow struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Content      string             `bson:"content, omitempty"`
	ShoppingList string             `bson:"shopping_list, omitempty"`
	Specs        *MenuSpecsRow      `bson:"specs, omitempty"`
	UserID       string             `bson:"userID"`
}

type MenuSpecsRow struct {
	MaxCalories string   `bson:"max_calories,omitempty"`
	MaxCarbs    string   `bson:"max_carbs,omitempty"`
	MaxProteins string   `bson:"max_proteins,omitempty"`
	MaxFats     string   `bson:"max_fats,omitempty"`
	Allergies   []string `bson:"allergies,omitempty"`
}

type MenuStore struct {
	client     *mongo.Client
	conf       MenuStoreConfig
	collection *mongo.Collection
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

	menusCollection := client.Database(conf.Database).Collection(conf.Collection)

	return &MenuStore{client: client, conf: conf, collection: menusCollection}, nil
}

func (s *MenuStore) Insert(ctx context.Context, row *MenuRow) (*MenuRow, error) {
	one, err := s.collection.InsertOne(ctx, row)

	if err != nil {
		return nil, errors.Wrapf(err, "inserting menu")
	}

	log.Printf("inserted menu with %s \n", one.InsertedID)

	return row, nil
}

func (s *MenuStore) GetByUserID(ctx context.Context, userID string) ([]*MenuRow, error) {
	var results []*MenuRow
	find, err := s.collection.Find(ctx, bson.E{Key: "userID", Value: userID})
	if err != nil {
		return nil, errors.Wrapf(err, "get menus by userID")
	}

	for find.Next(ctx) {
		var m MenuRow
		err := find.Decode(&m)
		if err != nil {
			log.Println(err)
		}
		results = append(results, &m)
	}

	return results, nil
}

func (s *MenuStore) Delete(ctx context.Context, menuID string) error {
	_, err := s.collection.DeleteOne(ctx, menuID)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("delete menu with id %s", menuID))
	}

	return nil
}
