package queue

import (
	"github.com/streadway/amqp"

	"github.com/mikarios/golib/queue/rabbitmq"
)

type RabbitMQConf struct {
	URL string
}

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewQueue(cfg *RabbitMQConf) (*RabbitMQ, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{Conn: conn, Ch: ch}, nil
}

func (r *RabbitMQ) Exchange() *rabbitmq.Exchange {
	return &rabbitmq.Exchange{Channel: r.Ch}
}

func (r *RabbitMQ) Queue() *rabbitmq.Queue {
	return &rabbitmq.Queue{Channel: r.Ch}
}
