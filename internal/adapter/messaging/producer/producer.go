package producer

import (
	"context"
	"encoding/json"

	"github.com/enesanbar/go-service/info"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/messaging/rabbitmq"
	"github.com/rabbitmq/amqp091-go"

	"go.uber.org/fx"
)

// TODO: Move this to a common package
type Producer struct {
	Logger  log.Factory
	Channel *rabbitmq.Channel
}

type ProducerParams struct {
	fx.In

	Logger  log.Factory
	Channel *rabbitmq.Channel `name:"default"`
}

func NewProducer(params ProducerParams) *Producer {
	return &Producer{
		Channel: params.Channel,
	}
}

func (p *Producer) Publish(ctx context.Context, messageName string, message any) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.Channel.Channel.PublishWithContext(
		ctx,
		info.ServiceName, // exchange
		messageName,      // routing key
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
