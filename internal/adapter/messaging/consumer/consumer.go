package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/messaging/messages"
	"github.com/enesanbar/go-service/messaging/rabbitmq"
	"github.com/enesanbar/go-service/wiring"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ConsumerParams struct {
	fx.In

	Logger          log.Factory
	Channel         *rabbitmq.Channel `name:"default"`
	Queue           *rabbitmq.Queue   `name:"url-shortener-worker"`
	MessageHandlers []MessageHandler  `group:"message-handlers"`
}

type RabbitMQQueueConsumer struct {
	logger          log.Factory
	Channel         *rabbitmq.Channel
	Queue           *rabbitmq.Queue
	Delivery        <-chan amqp091.Delivery
	MessageHandlers []MessageHandler
}

// New creates a pointer to the new instance of the HttpServer
func New(p ConsumerParams) (wiring.RunnableGroup, *RabbitMQQueueConsumer) {
	consumer := &RabbitMQQueueConsumer{
		logger:          p.Logger,
		Channel:         p.Channel,
		Queue:           p.Queue,
		MessageHandlers: p.MessageHandlers,
	}
	return wiring.RunnableGroup{
		Runnable: consumer,
	}, consumer
}

func (h *RabbitMQQueueConsumer) Start() error {
	// add recovery logic to the channel when channel/connection is closed
	msgs, err := h.Channel.Channel.Consume(
		h.Queue.Queue.Name, // queue
		"",                 // consumer
		true,               // auto ack
		false,              // exclusive
		false,              // no local
		false,              // no wait
		nil,                // args
	)
	if err != nil {
		return fmt.Errorf("error starting RabbitMQ consumer (%w)", err)
	}
	h.Delivery = msgs

	go func() {
		for d := range msgs {
			// TODO: throttle the message processing with worker pool
			go func(d amqp091.Delivery) {
				message := messages.Message[any]{}
				err := json.Unmarshal(d.Body, &message)
				if err != nil {
					h.logger.Bg().With(zap.Error(err)).Error("Failed to unmarshal message")
					return
				}

				// TODO: inject otel trace to tracing span
				for _, handler := range h.MessageHandlers {
					if handler.Name() == d.RoutingKey {
						err := handler.Handle(message)
						if err != nil {
							h.logger.Bg().With(zap.Error(err)).Error("Failed to handle message")
						}
						return
					}
				}
				h.logger.Bg().With(zap.String("messageName", d.RoutingKey)).Error("No handler found for message")
			}(d)
		}
		h.logger.Bg().Info("RabbitMQ consumer stopped")
	}()
	h.logger.Bg().Info("RabbitMQ consumer started")
	return nil
}

func (h *RabbitMQQueueConsumer) Stop() error {
	h.logger.Bg().Info("RabbitMQ consumer stopped")
	return nil
}
