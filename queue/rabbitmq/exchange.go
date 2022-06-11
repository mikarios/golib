package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Exchange struct {
	Channel    *amqp.Channel
	name       string
	kind       string
	durable    bool
	autoDelete bool
	internal   bool
	noWait     bool
	arguments  amqp.Table
}

func (e *Exchange) Name(v string) *Exchange {
	e.name = v

	return e
}

func (e *Exchange) Kind(v string) *Exchange {
	e.kind = v

	return e
}

func (e *Exchange) Durable(v bool) *Exchange {
	e.durable = v

	return e
}

func (e *Exchange) AutoDelete(v bool) *Exchange {
	e.autoDelete = v

	return e
}

func (e *Exchange) Internal(v bool) *Exchange {
	e.internal = v

	return e
}

func (e *Exchange) NoWait(v bool) *Exchange {
	e.noWait = v

	return e
}

func (e *Exchange) Arguments(v amqp.Table) *Exchange {
	e.arguments = v

	return e
}

func (e *Exchange) Declare() error {
	return e.Channel.ExchangeDeclare(
		e.name,
		e.kind,
		e.durable,
		e.autoDelete,
		e.internal,
		e.noWait,
		e.arguments,
	)
}

func (e *Exchange) DeclarePassive() error {
	return e.Channel.ExchangeDeclare(
		e.name,
		e.kind,
		e.durable,
		e.autoDelete,
		e.internal,
		e.noWait,
		e.arguments,
	)
}
