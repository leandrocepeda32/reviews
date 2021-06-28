package rest

import (
	"github.com/go-chi/chi"
)

func RegisterReviewsRoutes(r chi.Router, handler *ReviewHandler) {
	r.Use(securityMiddleware)
	r.Route("/reviews", func(r chi.Router) {
		r.Post("/", handler.createReview)
		r.Delete("/{id}", handler.DeleteReview)
		r.Get("/article/{id}", handler.GetArticleReviews)
		r.Get("/article/rating/{id}", handler.GetArticleRating)
	})
}