package review

import (
	"context"
	"time"

	"github.com/leandrocepeda32/reviews/internal/domain/review/vo"
	"github.com/leandrocepeda32/reviews/internal/utils/errors"
	uuid "github.com/satori/go.uuid"
)

type Service interface {
	CreateReview(ctx context.Context, createReview *vo.CreateReview) error
	GetArticleReviews(ctx context.Context, id string) (*[]Review, error)
}

type Repository interface {
	Save(ctx context.Context, review *Review) error
	getAllReviewsByArticle(ctx context.Context, id string) (*[]Review, error)
}

type reviewService struct {
	repository Repository
}

func NewReviewService(repository Repository) *reviewService {
	return &reviewService {
		repository: repository,
	}
}


// Al momento de crear una review, calculo el nuevo rating
func (rs *reviewService) CreateReview(ctx context.Context, createReview *vo.CreateReview) error {
	err := createReview.Validate()

	if err != nil {
		return err
	}

	review := Review{
		Id: uuid.NewV4(),
		Comment: createReview.Comment,
		Score: createReview.Score,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ArticleId: createReview.ArticleId,
	}

	return rs.repository.Save(ctx, &review)
}

func (rs *reviewService) GetArticleReviews(ctx context.Context, id string) (*[]Review, error) {
	reviews, err := rs.repository.getAllReviewsByArticle(ctx, id)

	if err != nil {
		if errCast, ok := err.(*errors.RestClientError); ok {
			if errCast.StatusCode == 404 {
				return nil, errors.NotFound
			}
		}
		return nil, err
	}

	return reviews, nil
}



