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

	// Define a filter to match documents by userID
	filter := bson.M{"userID": userID}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrapf(err, "get menus by userID")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var m MenuRow
		if err := cursor.Decode(&m); err != nil {
			log.Println(err)
			continue // Skip current iteration on error
		}
		results = append(results, &m)
	}

	log.Printf("results: %v \n", results)

	return results, nil
}

func (s *MenuStore) Delete(ctx context.Context, menuID string) error {
	// Convert menuID string to a BSON ObjectID
	objectID, err := primitive.ObjectIDFromHex(menuID)
	if err != nil {
		return errors.Wrap(err, "invalid ObjectID format")
	}

	// Construct a filter to match the menu by its ObjectID
	filter := bson.M{"_id": objectID}

	_, err = s.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("delete menu with id %s", menuID))
	}

	return nil
}
