package review

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Comment string `json:"comment"`
	Score int `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ArticleId string `json:"article_id"`
	UserId string `json:"user_id"`
}