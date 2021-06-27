package rest

import (
	"github.com/go-chi/chi"
)

func RegisterReviewsRoutes(r chi.Router, handler *ReviewHandler) {
	r.Use(securityMiddleware)
	r.Route("/reviews", func(r chi.Router) {
		r.Get("/{id}", handler.GetArticleReviews)
		r.Post("/", handler.createReview)
	})
}