package dbmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collection = "person"

type CreatePersonParams struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Age  int    `json:"age,omitempty" bson:"age,omitempty"`
	Bio  string `json:"bio,omitempty" bson:"bio,omitempty"`
}

func (store *MongoStore) InsertPerson(ctx context.Context, args *CreatePersonParams) (*Person, error) {
	res, err := store.db.Collection(collection).InsertOne(ctx, args)
	if err != nil {
		return nil, err
	}

	person := &Person{
		ID:   res.InsertedID.(primitive.ObjectID),
		Name: args.Name,
		Age:  args.Age,
		Bio:  args.Bio,
	}

	return person, nil
}

func (store *MongoStore) GetListPerson(ctx context.Context, page int64, perPage int64) ([]*Person, error) {
	cursor, err := store.db.Collection(collection).Find(ctx, map[string]any{}, options.Find().SetLimit(perPage).SetSkip((page-1)*perPage))
	if err != nil {
		return nil, err
	}

	people := []*Person{}
	err = cursor.All(ctx, &people)

	return people, nil
}
