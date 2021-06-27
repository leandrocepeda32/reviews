package review

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Review struct {
	Id uuid.UUID `json:"id"`
	Comment string `json:"comment"`
	Score int `json:"score"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ArticleId string `json:"article_id"`
	UserId string `json:"user_id"`
}