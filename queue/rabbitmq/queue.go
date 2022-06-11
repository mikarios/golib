package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Queue struct {
	Channel          *amqp.Channel
	name             string
	durable          bool
	deleteWhenUnused bool
	exclusive        bool
	noWait           bool
	arguments        amqp.Table
}

func (q *Queue) Name(v string) *Queue {
	q.name = v

	return q
}

func (q *Queue) Durable(v bool) *Queue {
	q.durable = v

	return q
}

func (q *Queue) DeleteWhenUnused(v bool) *Queue {
	q.deleteWhenUnused = v

	return q
}

func (q *Queue) Exclusive(v bool) *Queue {
	q.exclusive = v

	return q
}

func (q *Queue) NoWait(v bool) *Queue {
	q.noWait = v

	return q
}

func (q *Queue) Arguments(v amqp.Table) *Queue {
	q.arguments = v

	return q
}

func (q *Queue) Declare() (amqp.Queue, error) {
	return q.Channel.QueueDeclare(
		q.name,
		q.durable,
		q.deleteWhenUnused,
		q.exclusive,
		q.noWait,
		q.arguments,
	)
}
