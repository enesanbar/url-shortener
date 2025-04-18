package events

import (
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/messaging/consumer"
	"github.com/enesanbar/go-service/messaging/messages"

	"go.uber.org/zap"
)

type MappingUpdatedHandler struct {
	logger log.Factory
}

func NewMappingUpdatedEventHandler(logger log.Factory) *MappingUpdatedHandler {
	return &MappingUpdatedHandler{
		logger: logger,
	}
}
func (h *MappingUpdatedHandler) Handle(message messages.Message[any]) error {
	h.logger.Bg().With(zap.Any("message", message)).Info("message received")
	// Invalidate cache
	return nil
}

func (h *MappingUpdatedHandler) Properties() consumer.MessageProperties {
	return consumer.MessageProperties{
		QueueName:   "url-shortener-worker",
		MessageName: "mappingUpdated",
	}
}
