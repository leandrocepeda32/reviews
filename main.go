package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/leandrocepeda32/reviews/internal/domain/review"
	"github.com/leandrocepeda32/reviews/internal/rest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	
	//MONGO
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal()
	}

	collection := client.Database("review").Collection("review")

	//REPOSITORIES
	reviewRepository := review.NewReviewRepository(collection)

	//SERVICES
	reviewService := review.NewReviewService(reviewRepository)

	//HANDLERS
	reviewHandler := rest.NewReviewHandler(reviewService)

	//START
	router := chi.NewRouter()
	rest.RegisterReviewsRoutes(router, reviewHandler)

	log.Print("Availability routes")
	for _, a := range router.Routes() {
		for _, b := range a.SubRoutes.Routes() {
			log.Print(fmt.Sprint(strings.ReplaceAll(a.Pattern, "/*", ""), b.Pattern))
		}

	}
	port := ":8020"
	log.Print(fmt.Sprint("Starting server at port", port))
	log.Fatal(http.ListenAndServe(port, router))





}