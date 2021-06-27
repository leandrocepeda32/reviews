package rating

import (
	"context"
	"log"

	"github.com/leandrocepeda32/reviews/internal/domain/rating/vo"
	uuid "github.com/satori/go.uuid"
)

type Service interface {
	CalculateNewRating(ctx context.Context, createRating *vo.CreateRating) error
	GetArticleRating(ctx context.Context, id string) (float64, error)
	GetAccumulatedRating(ctx context.Context, id string) (vo.AccumulatedRating, error)
}

type Repository interface {
	Save(ctx context.Context, rating *Rating) error
	GetArticleRating(ctx context.Context, id string) (*[]Rating, error)
}

type ratingService struct {
	repository Repository
}

func NewRatingService(repository Repository) Service {
	return &ratingService {
		repository: repository,
	}
}

func (rs *ratingService) CalculateNewRating(ctx context.Context, createRating *vo.CreateRating) error {

	articleId := createRating.ArticleId
	accumulatedRating,err := rs.GetAccumulatedRating(ctx, articleId)

	if err != nil {
		return err
	}

	newScore := accumulatedRating.LastScoreSum + float64(createRating.Score)
	newCount := accumulatedRating.LastScoreCount + 1

	newRating := Rating{
		Id: uuid.NewV4(),
		IdArticle: articleId,
		ScoreSum: newScore,
		ScoreCount: newCount,
	}

	err = rs.repository.Save(ctx, &newRating)

	if err != nil {
		return err
	}

	return nil

}

func (rs *ratingService) GetArticleRating(ctx context.Context, id string) (float64, error) {
	accumulatedRating, err := rs.GetAccumulatedRating(ctx, id)

	if err != nil {
		return 0,err
	}

	return accumulatedRating.LastScoreSum / float64(accumulatedRating.LastScoreCount), nil

}

func(rs *ratingService) GetAccumulatedRating(ctx context.Context, id string) (vo.AccumulatedRating, error) {
	ratings, err := rs.repository.GetArticleRating(ctx, id)

	if err != nil {
		log.Fatal(err)
	}
	
	var lastScore float64 = 0
	var lastCount int64 = 0

	for _,rating := range *ratings {
		lastScore += rating.ScoreSum
		lastCount += 1
	}

	accumulatedRating := vo.AccumulatedRating{
		LastScoreSum: lastScore,
		LastScoreCount: lastCount,
	}

	return accumulatedRating, nil

}

