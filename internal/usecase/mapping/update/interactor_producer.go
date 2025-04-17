package update

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
func NewUpdateMappingInteractorProducer(p ProducerParams) Service {
	return &InteractorProducer{
		logger:      p.Logger,
		producer:    p.Producer,
		next:        p.Next,
		messageName: "MappingUpdated",
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
	err = a.producer.Publish(ctx, a.messageName, messages.Message[MappingUpdatedPayload]{
		Metadata: messages.Metadata{
			MessageName:   a.messageName,
			PublishDate:   time.Now().UTC(),
			PublisherName: info.ServiceName,
		},
		Payload: MappingUpdatedPayload{
			ID:        result.ID,
			Code:      result.Code,
			URL:       result.URL,
			ExpiresAt: result.ExpiresAt,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		},
	})

	if err != nil {
		a.logger.For(ctx).Error("failed to publish message", zap.Error(err))
	}

	return result, nil
}
