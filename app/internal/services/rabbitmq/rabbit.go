package rabbitmq

import (
	"context"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	rabbit "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	connection rabbit.Channel
	queue      rabbit.Queue
	ctx        context.Context
}

func NewRabbitMQ(conn amqp091.Channel) *RabbitMQ {
	return &RabbitMQ{
		connection: conn,
	}
}

func (r *RabbitMQ) Produce(body []byte) error {
	err := r.connection.PublishWithContext(r.ctx,
		"",           // exchange
		r.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		rabbit.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		return err
	}

	log.Printf(" [x] Queue message sent %s\n", body)

	return nil
}

func (r *RabbitMQ) Consume() (interface{}, error) {
	return nil, nil
}

func (r *RabbitMQ) DeclareQueue(queue string) error {
	var err error

	r.queue, err = r.connection.QueueDeclarePassive(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	r.ctx = ctx
	defer cancel()

	log.Printf(" [x] Queue declared %s\n", queue)

	return nil
}
