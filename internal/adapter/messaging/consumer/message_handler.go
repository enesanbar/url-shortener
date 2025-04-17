package consumer

import "github.com/enesanbar/go-service/messaging/messages"

// TODO: move this to the go-service package
type MessageHandler interface {
	Handle(message messages.Message[any]) error
	Name() string
}
