package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"session-service/config"
	"session-service/database"
	"session-service/models"
	"time"

	"github.com/streadway/amqp"
)

func ConsumeSessionEvents() {
	conn, err := amqp.Dial(config.RabbitMQURL)
	if err != nil {
		log.Fatal("RabbitMQ connection error:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	q, err := ch.QueueDeclare(config.SessionQueue, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range msgs {
			var session models.Session
			err := json.Unmarshal(msg.Body, &session)
			if err != nil {
				log.Println("Failed to decode session message:", err)
				continue
			}

			session.Timestamp = time.Now()
			collection := database.GetCollection()
			_, err = collection.InsertOne(context.TODO(), session)
			if err != nil {
				log.Println("Failed to save session:", err)
			} else {
				log.Println("Session saved for user:", session.Email)
			}
		}
	}()
}
