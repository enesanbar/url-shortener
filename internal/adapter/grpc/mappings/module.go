package mappings

import "go.uber.org/fx"

var Module = fx.Module(
	"mappings",
	fx.Provide(NewGRPCServer),    // creates a new GRPC server. move to go-service
	fx.Invoke(NewMappingsServer), // registers the mappings server with the GRPC server
)
