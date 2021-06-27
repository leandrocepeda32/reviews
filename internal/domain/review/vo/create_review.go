package vo

import (
	"time"

	"github.com/leandrocepeda32/reviews/internal/domain"
	"github.com/leandrocepeda32/reviews/internal/utils/errors"
)

type CreateReview struct {
	Comment string `json:"comment"`
	Score int64 `json:"score"`
	ArticleId string `json:"article_id"`
}

func (cr *CreateReview) Validate() error {

	restClient := domain.NewRestClient(5 * time.Second)

	err := restClient.GetArticle(cr.ArticleId)

	if err != nil {
		return errors.NewBusinessError("invalid article")
	}
	
	if len(cr.Comment) < 10 {
		return errors.NewBusinessError("invalid comment")
	}

	if(cr.Score < 1 || cr.Score > 5) {
		return errors.NewBusinessError("invalid score")
	}

	return nil
}