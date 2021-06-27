package review

import (
	"context"
	"time"

	"github.com/leandrocepeda32/reviews/internal/domain/review/vo"
	"github.com/leandrocepeda32/reviews/internal/utils"
	"github.com/leandrocepeda32/reviews/internal/utils/errors"
	uuid "github.com/satori/go.uuid"
)

type Service interface {
	CreateReview(ctx context.Context, createReview *vo.CreateReview) error
	GetArticleReviews(ctx context.Context, id string) (*[]Review, error)
	GetArticleRating(ctx context.Context, id string) (*vo.ReviewsRating, error)
}

type Repository interface {
	Save(ctx context.Context, review *Review) error
	GetAllReviewsByArticle(ctx context.Context, id string) (*[]Review, error)
}

type ReviewMessageBroker interface {
	Created(ctx context.Context, review Review)
}

type reviewService struct {
	repository Repository
	messageBroker ReviewMessageBroker
}

func NewReviewService(repository Repository, messageBroker ReviewMessageBroker) Service {
	return &reviewService {
		repository: repository,
		messageBroker: messageBroker,
	}
}


func (rs *reviewService) CreateReview(ctx context.Context, createReview *vo.CreateReview) error {
	err := createReview.Validate()

	if err != nil {
		return err
	}

	userId := ctx.Value("user").(utils.ContextValues).Get("user_id")
	
	review := Review{
		Id: uuid.NewV4(),
		Comment: createReview.Comment,
		Score: createReview.Score,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ArticleId: createReview.ArticleId,
		UserId: userId,
	}

	err = rs.repository.Save(ctx, &review)

	if err != nil {
		return err
	}

	rs.messageBroker.Created(ctx, review)

	return nil
}

func (rs *reviewService) GetArticleReviews(ctx context.Context, id string) (*[]Review, error) {
	reviews, err := rs.repository.GetAllReviewsByArticle(ctx, id)

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

func(rs *reviewService) GetArticleRating(ctx context.Context, id string) (*vo.ReviewsRating, error) {
	reviews, err := rs.repository.GetAllReviewsByArticle(ctx, id)

	if err != nil {
		if errCast, ok := err.(*errors.RestClientError); ok {
			if errCast.StatusCode == 404 {
				return nil, errors.NotFound
			}
		}
		return nil, err
	}

	if len(*reviews) == 0 {
		return nil, errors.NotFound
	}

	var scoreSum int = 0;

	for _, review := range(*reviews) {
		scoreSum += review.Score
	}

	rating := float64(scoreSum) / float64(len(*reviews))

	reviewsRating := vo.ReviewsRating{
		ArticleId: id,
		Rating: rating,
	}

	return &reviewsRating, nil
}