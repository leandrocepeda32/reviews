package rest

import (
	"github.com/go-chi/chi"
)

func RegisterReviewsRoutes(r chi.Router, handler *ReviewHandler) {
	r.Use(authMiddleware)
	r.Route("/reviews", func(r chi.Router) {
		r.Post("/", handler.createReview)
		r.Delete("/{id}", handler.DeleteReview)
	})
	r.Route("/articles", func(r chi.Router) {
		r.Get("/{id}/reviews", handler.GetArticleReviews)
		r.Get("/{id}/rating", handler.GetArticleRating)
	})
}