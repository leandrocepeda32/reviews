package review

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type repositoryMongo struct {
	Collection *mongo.Collection
}

func NewReviewRepository(collection *mongo.Collection) Repository {
	return &repositoryMongo{Collection: collection}
}

func (r *repositoryMongo) Save(ctx context.Context, review *Review) error {
	_, err := r.Collection.InsertOne(ctx, review)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryMongo) GetAllReviewsByArticle(ctx context.Context, id string) (*[]Review, error) {
	filter := bson.M{"articleid" : id}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	
	var reviews []Review
	if err = cursor.All(ctx, &reviews); err != nil {
		log.Fatal(err)
	}

	return &reviews, nil

}
