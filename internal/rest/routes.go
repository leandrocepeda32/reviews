package rest

import (
	"github.com/go-chi/chi"
	_ "github.com/leandrocepeda32/reviews/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterReviewsRoutes(r chi.Router, handler *ReviewHandler) {
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)
		r.Route("/reviews", func(r chi.Router) {
			r.Post("/", handler.CreateReview)
			r.Delete("/{id}", handler.DeleteReview)
		})
		r.Route("/articles", func(r chi.Router) {
			r.Get("/{id}/reviews", handler.GetArticleReviews)
			r.Get("/{id}/rating", handler.GetArticleRating)
		})
	})
	
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	
}