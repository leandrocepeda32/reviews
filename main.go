package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/leandrocepeda32/reviews/internal/domain/review"
	"github.com/leandrocepeda32/reviews/internal/rabbitmq"
	"github.com/leandrocepeda32/reviews/internal/rest"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Reviews Microservice
// @version 1.0
// @description Microservice to add reviews to articles
// @BasePath /
func main() {

	//RABBIT
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("error connecting rabbit")
	}

	ch, err := conn.Channel()
	if err != nil {
		panic("error open rabbit channel")
	}

	reviewMessageBroker, err := rabbitmq.NewReview(ch)

	if err != nil {
		panic("error initializing review message broker")
	}
	
	//MONGO
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("review").Collection("review")

	//REPOSITORIES
	reviewRepository := review.NewReviewRepository(collection)

	//SERVICES
	reviewService := review.NewReviewService(reviewRepository, reviewMessageBroker)

	//HANDLERS
	reviewHandler := rest.NewReviewHandler(reviewService)

	//ROUTES
	router := chi.NewRouter()


	rest.RegisterReviewsRoutes(router, reviewHandler)
	
	//START
	port := ":8020"
	log.Print(fmt.Sprint("Starting server at port", port))
	log.Fatal(http.ListenAndServe(port, router))
}