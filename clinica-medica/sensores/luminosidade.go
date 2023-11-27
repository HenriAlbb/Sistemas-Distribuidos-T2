package main

import (
	"context"
	"encoding/json"
	"example/m/entities/sensores"
	"log"
	"math/rand"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
/*
	q, err := ch.QueueDeclare(
		"fila-1",
		false,
		false,
		false,
		false,
		nil,
	)*/
	err = ch.ExchangeDeclare(
		"fila-1",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
)
	failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for {

		luminosidade := sensores.Luminosidade{"luminosidade-1", "luminosidade", "", "", "", rand.Float64() * 1000}
		msg, _ := json.Marshal(luminosidade)

		//body := "Hello World!"
		err = ch.PublishWithContext(ctx,
			"fila-1",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        msg,
			})

		time.Sleep(2 * time.Second)

		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", msg)

	}

}
