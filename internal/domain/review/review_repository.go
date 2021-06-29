package review

import (
	"context"
	"fmt"
	"log"

	"github.com/leandrocepeda32/reviews/internal/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		log.Print(err)
		return nil, err
	}
	
	var reviews []Review
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	if(len(reviews) == 0) {
		return nil, errors.NotFound
	}

	return &reviews, nil

}

func (r *repositoryMongo) Delete(ctx context.Context, id string) error {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	
	if err != nil {
		return errors.NewBusinessError("Invalid Id")
	}

	filter := bson.M{"_id" : idPrimitive}

	deleteResult, err := r.Collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if deleteResult.DeletedCount == 0 {
		return errors.NewBusinessError(fmt.Sprintf("The id %s doesn't exist", id))
	}

	return nil
}
