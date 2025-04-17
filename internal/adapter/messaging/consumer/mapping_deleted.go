package consumer

import (
	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/messaging/messages"
	"go.uber.org/zap"
)

type MappingDeletedHandler struct {
	logger log.Factory
}

func NewMappingDeletedEventHandler(logger log.Factory) *MappingDeletedHandler {
	return &MappingDeletedHandler{
		logger: logger,
	}
}
func (h *MappingDeletedHandler) Handle(message messages.Message[any]) error {
	h.logger.Bg().With(zap.Any("message", message)).Info("message received")
	return nil
}

func (h *MappingDeletedHandler) Name() string {
	return "MappingDeleted"
}
