package rating

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryMongo struct {
	Collection *mongo.Collection
}

func NewRatingRepository(collection *mongo.Collection) Repository {
	return &repositoryMongo{Collection: collection}
}

func (r *repositoryMongo) Save(ctx context.Context, rating *Rating) error {
	_, err := r.Collection.InsertOne(ctx, rating)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryMongo) GetArticleRating(ctx context.Context, id string) (*[]Rating, error) {
	filter := bson.M{"articleid" : id}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	
	var ratings []Rating
	if err = cursor.All(ctx, &ratings); err != nil {
		log.Fatal(err)
	}

	return &ratings, nil
}