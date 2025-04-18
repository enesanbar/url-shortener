package events

import (
	"github.com/enesanbar/go-service/messaging/consumer"
	"github.com/enesanbar/go-service/messaging/rabbitmq"
	"github.com/enesanbar/go-service/wiring"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"messaging",

	// Register the event handlers
	fx.Provide(
		consumer.AsMessageHandler(NewMappingCreatedEventHandler),
		consumer.AsMessageHandler(NewMappingUpdatedEventHandler),
		consumer.AsMessageHandler(NewMappingDeletedEventHandler),
	),

	// run the consumer
	// TODO: Later find a way to provide this to FX automatically in go-service
	fx.Provide(func(p consumer.RabbitMQConsumerParams) (wiring.RunnableGroup, error) {
		return consumer.RabbitMQConsumerFactory(consumer.RabbitMQConsumersParams{
			QueueName:       "url-shortener-worker",
			Queues:          p.Queues,
			Channels:        p.Channels,
			Logger:          p.Logger,
			MessageHandlers: p.MessageHandlers,
		})
	}),

	// Register connections as Connection, so that they can recover from failures
	// TODO: Later find a way to provide this to FX automatically in go-service
	fx.Provide(
		fx.Annotated{
			Group: "connections",
			Target: func(connections map[string]*rabbitmq.Connection) wiring.Connection {
				return connections["default"]
			},
		},
		fx.Annotated{
			Group: "connections",
			Target: func(channels map[string]*rabbitmq.Channel) wiring.Connection {
				return channels["default"]
			},
		},
	),
)
