package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/leandrocepeda32/reviews/internal/domain/review"
	"github.com/leandrocepeda32/reviews/internal/domain/review/vo"
	"github.com/leandrocepeda32/reviews/internal/utils/errors"
)

type Service interface {
	CreateReview(ctx context.Context, createReview *vo.CreateReview) error
	GetArticleReviews(ctx context.Context, id string) (*[]review.Review, error)
}

type ReviewHandler struct {
	service Service
}

func NewReviewHandler(service Service) *ReviewHandler {
	return &ReviewHandler{
		service: service,
	}
}



func (rh *ReviewHandler) createReview(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	createReview := &vo.CreateReview{}
	err := json.NewDecoder(r.Body).Decode(createReview)

	if err != nil {
		ErrorResponse(w, 400, err.Error())
		return 
	}

	err = rh.service.CreateReview(ctx, createReview)
	
	if err != nil {
		ErrorResponse(w, 400, err.Error())
		return
	}

	WebResponse(w, 201, createReview)
}

func (rh *ReviewHandler) GetArticleReviews(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		ErrorResponse(w, 400, "id is required")
		return
	}
	ctx := r.Context()

	reviews, err := rh.service.GetArticleReviews(ctx, id)

	if err != nil {
		if err == errors.NotFound {
			ErrorResponse(w, 404, "not_found")
		}
		ErrorResponse(w, 503, err.Error())
	}

	if len(*reviews) == 0 {
		ErrorResponse(w, 404, "The article has not reviews")
		return
	}

	WebResponse(w, 200, reviews)
}
