package create

import (
	"context"
	"time"

	"github.com/enesanbar/go-service/info"
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/messaging/messages"
	"github.com/enesanbar/url-shortener/internal/adapter/messaging/producer"
	"github.com/enesanbar/url-shortener/internal/domain"
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
func NewCreateMappingInteractorProducer(p ProducerParams) Service {
	return &InteractorProducer{
		logger:      p.Logger,
		producer:    p.Producer,
		next:        p.Next,
		messageName: "MappingCreated",
	}
}

type InteractorProducer struct {
	logger      log.Factory
	producer    *producer.Producer
	next        Service
	messageName string
}

// Execute orchestrates the use case
func (a InteractorProducer) Execute(ctx context.Context, input *Request) (*domain.Mapping, error) {
	result, err := a.next.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	// publish message
	err = a.producer.Publish(ctx, a.messageName, messages.Message[domain.MappingCreatedPayload]{
		Metadata: messages.Metadata{
			MessageName:   a.messageName,
			PublishDate:   time.Now().UTC(),
			PublisherName: info.ServiceName,
		},
		Payload: domain.MappingCreatedPayload{
			MappingID:   result.ID,
			ShortURL:    result.Code,
			OriginalURL: result.URL,
		},
	})

	if err != nil {
		a.logger.For(ctx).Error("failed to publish message", zap.Error(err))
	}

	return result, nil
}
