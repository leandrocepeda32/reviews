package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/leandrocepeda32/reviews/internal/domain/review"
	"github.com/streadway/amqp"
)

type Review struct {
	ch *amqp.Channel
}

func NewReview(channel *amqp.Channel) (*Review, error) {
	return &Review{
		ch: channel,
	}, nil
}

// Created publishes a message indicating a review was created.
func (r *Review) Created(ctx context.Context, review review.Review) {
	r.publish(ctx, review)
}

func (t *Review) publish(ctx context.Context, e interface{}) {
	
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(e); err != nil {
		failOnError(err, "Failed encoding data")
	}


	q, err := t.ch.QueueDeclare(
		"create_review", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare a queue")

	err = t.ch.Publish(
		"",    // exchange
		q.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        b.Bytes(),
			Timestamp:   time.Now(),
		})
	
	log.Printf(" [x] Sent %s", b.Bytes())
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}