package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/leandrocepeda32/reviews/internal/domain/review"
	"github.com/leandrocepeda32/reviews/internal/domain/review/vo"
	"github.com/leandrocepeda32/reviews/internal/utils/errors"
)

type Service interface {
	CreateReview(ctx context.Context, createReview *vo.CreateReview) error
	GetArticleReviews(ctx context.Context, id string) (*[]review.Review, error)
	GetArticleRating(ctx context.Context, id string) (*vo.ReviewsRating, error)
	DeleteReview(ctx context.Context, id string) error
}

type ReviewHandler struct {
	service Service
}

func NewReviewHandler(service Service) *ReviewHandler {
	return &ReviewHandler{
		service: service,
	}
}

// CreateReview godoc
// @Summary Add a review
// @Description Add a review for an article
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param review body vo.CreateReview true "Add review"
// @Success 200 {object} vo.CreateReview
// @Failure 400 {object} errors.ErrCustom
// @Failure 404 {object} errors.ErrCustom
// @Router /reviews [post]
func (rh *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	createReview := &vo.CreateReview{}
	err := json.NewDecoder(r.Body).Decode(createReview)

	if err != nil {
		ErrorResponse(w, 400, "Can't unmarshal JSON object into struct")
		return
	}

	restClient := NewRestClient(5 * time.Second)
	err = restClient.GetArticle(createReview.ArticleId)

	if err != nil {
		ErrorResponse(w, 400, fmt.Sprintf("The article with id %s doesn't exist", createReview.ArticleId))
		return
	}

	err = rh.service.CreateReview(ctx, createReview)
	
	if err != nil {
		ErrorResponse(w, 400, err.Error())
		return
	}

	WebResponse(w, 201, createReview)
}

// GetArticleReviews godoc
// @Summary List of reviews for an article
// @Description Get the reviews of an article
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param id path string true "article id"
// @Success 200 {array} []review.Review
// @Failure 400 {object} errors.ErrCustom
// @Failure 404 {object} errors.ErrCustom
// @Router /articles/{id}/reviews [get]
func (rh *ReviewHandler) GetArticleReviews(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		ErrorResponse(w, 400, "id is required")
		return
	}

	restClient := NewRestClient(5 * time.Second)
	err := restClient.GetArticle(id)

	if err != nil {
		ErrorResponse(w, 400, fmt.Sprintf("The article with id %s doesn't exist", id))
		return
	}

	ctx := r.Context()

	reviews, err := rh.service.GetArticleReviews(ctx, id)

	if err != nil {
		if err == errors.NotFound {
			ErrorResponse(w, 404, "The article has no reviews")
			return
		}
		ErrorResponse(w, 503, err.Error())
		return
	}

	WebResponse(w, 200, reviews)
}

// GetArticleRating godoc
// @Summary Rating for an article
// @Description Get the rating of an article
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param id path string true "article id"
// @Success 200 {array} []vo.ReviewsRating
// @Failure 400 {object} errors.ErrCustom
// @Failure 404 {object} errors.ErrCustom
// @Router /articles/{id}/rating [get]
func (rh *ReviewHandler) GetArticleRating(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		ErrorResponse(w, 400, "Id is required")
		return
	}

	restClient := NewRestClient(5 * time.Second)
	err := restClient.GetArticle(id)

	if err != nil {
		ErrorResponse(w, 400, fmt.Sprintf("The article with id %s doesn't exist", id))
		return
	}

	ctx := r.Context()

	articleRating, err := rh.service.GetArticleRating(ctx, id)

	if err != nil {
		if err == errors.NotFound {
			ErrorResponse(w, 404, "The article has no reviews")
			return
		}
		ErrorResponse(w, 503, err.Error())
		return
	}


	WebResponse(w, 200, articleRating)
}


// DeleteReview godoc
// @Summary Delete a review
// @Description Delete a review by id
// @Tags reviews
// @Accept  json
// @Produce  json
// @Param  id path string true "review id"
// @Success 200
// @Failure 400 {object} errors.ErrCustom
// @Failure 404 {object} errors.ErrCustom
// @Router /reviews/{id} [delete]
func (rh *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		ErrorResponse(w, 400, "Id is required")
		return
	}
	ctx := r.Context()

	err := rh.service.DeleteReview(ctx, id)

	if err != nil {
		ErrorResponse(w, 400, err.Error())
		return
	}

	w.WriteHeader(200)
}