package messaging

import (
	"github.com/enesanbar/go-service/info"
	"github.com/enesanbar/go-service/messaging/rabbitmq"
	"github.com/enesanbar/go-service/wiring"
	"github.com/enesanbar/url-shortener/internal/adapter/messaging/consumer"
	"github.com/enesanbar/url-shortener/internal/adapter/messaging/producer"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"messaging",
	fx.Provide(
		fx.Annotated{
			Name: "default",
			Target: func(connetions map[string]*rabbitmq.Connection) *rabbitmq.Connection {
				return connetions["default"]
			},
		},
		fx.Annotated{
			Group: "connections",
			Target: func(connetions map[string]*rabbitmq.Connection) wiring.Connection {
				return connetions["default"]
			},
		},
		fx.Annotated{
			Name: "default",
			Target: func(channels map[string]*rabbitmq.Channel) *rabbitmq.Channel {
				return channels["default"]
			},
		},
		fx.Annotated{
			Group: "connections",
			Target: func(channels map[string]*rabbitmq.Channel) wiring.Connection {
				return channels["default"]
			},
		},
		fx.Annotated{
			Name: "url-shortener-worker",
			Target: func(queues map[string]*rabbitmq.Queue) *rabbitmq.Queue {
				return queues["url-shortener-worker"]
			},
		},

		// we are invoking exchanges just to trigger the creation of the exchange
		fx.Annotated{
			Name: info.ServiceName,
			Target: func(exchanges map[string]*rabbitmq.Exchange) *rabbitmq.Exchange {
				return exchanges[info.ServiceName]
			},
		},
	),

	fx.Provide(producer.NewProducer),
	fx.Provide(consumer.New),
	fx.Provide(
		fx.Annotate(
			consumer.NewMappingCreatedEventHandler,
			fx.As(new(consumer.MessageHandler)),
			fx.ResultTags(`group:"message-handlers"`),
		),
		fx.Annotate(
			consumer.NewMappingUpdatedEventHandler,
			fx.As(new(consumer.MessageHandler)),
			fx.ResultTags(`group:"message-handlers"`),
		),
		fx.Annotate(
			consumer.NewMappingDeletedEventHandler,
			fx.As(new(consumer.MessageHandler)),
			fx.ResultTags(`group:"message-handlers"`),
		),
	),
)
