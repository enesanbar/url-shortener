package deletion

import (
	"context"
	"time"

	"github.com/enesanbar/go-service/info"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/messaging/messages"
	"github.com/enesanbar/url-shortener/internal/adapter/messaging/producer"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ProducerParams struct {
	fx.In

	Logger   log.Factory
	Producer *producer.Producer
	Next     Service `name:"interactor"`
}

// NewCreateMappingInteractorProducer creates new InteractorProducer with its dependencies
func NewDeleteMappingInteractorProducer(p ProducerParams) Service {
	return &InteractorProducer{
		logger:      p.Logger,
		producer:    p.Producer,
		next:        p.Next,
		messageName: "MappingDeleted",
	}
}

type InteractorProducer struct {
	logger      log.Factory
	producer    *producer.Producer
	next        Service
	messageName string
}

// Execute orchestrates the use case
func (a InteractorProducer) Execute(ctx context.Context, code string) error {
	err := a.next.Execute(ctx, code)
	if err != nil {
		return err
	}

	// publish message
	err = a.producer.Publish(ctx, a.messageName, messages.Message[MappingDeletedPayload]{
		Metadata: messages.Metadata{
			MessageName:   a.messageName,
			PublishDate:   time.Now().UTC(),
			PublisherName: info.ServiceName,
		},
		Payload: MappingDeletedPayload{
			Code: code,
		},
	})

	if err != nil {
		a.logger.For(ctx).Error("failed to publish message", zap.Error(err))
	}

	return nil
}
