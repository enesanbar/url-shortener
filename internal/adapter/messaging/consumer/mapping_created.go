package consumer

import (
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/messaging/messages"
	"go.uber.org/zap"
)

type MappingCreatedHandler struct {
	logger log.Factory
}

func NewMappingCreatedEventHandler(logger log.Factory) *MappingCreatedHandler {
	return &MappingCreatedHandler{
		logger: logger,
	}
}
func (h *MappingCreatedHandler) Handle(message messages.Message[any]) error {
	h.logger.Bg().With(zap.Any("message", message)).Info("message received")
	// payload := message.Payload.(domain.MappingCreatedPayload)
	return nil
}

func (h *MappingCreatedHandler) Name() string {
	return "MappingCreated"
}
