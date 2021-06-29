package vo

import (
	"github.com/leandrocepeda32/reviews/internal/utils/errors"
)

type CreateReview struct {
	ArticleId string `json:"article_id"`
	Comment string `json:"comment"`
	Score int `json:"score"`
}

func (cr *CreateReview) Validate() error {
	
	if len(cr.Comment) < 10 {
		return errors.NewBusinessError("Comment must have at least 10 characters")
	}

	if(cr.Score < 1 || cr.Score > 5) {
		return errors.NewBusinessError("Score must be between 1 and 5")
	}

	return nil
}