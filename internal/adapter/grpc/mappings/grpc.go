package mappings

import (
	"fmt"
	"net"

	"github.com/enesanbar/go-service/log"
	"github.com/enesanbar/go-service/wiring"
	"go.uber.org/fx"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

type GRPCServerParams struct {
	fx.In

	Logger log.Factory
}

type GRPCServer struct {
	grpcServer *grpc.Server
	logger     log.Factory
}

func NewGRPCServer(p GRPCServerParams) (wiring.RunnableGroup, *GRPCServer) {
	s := grpc.NewServer()

	grpcServer := &GRPCServer{
		grpcServer: s,
		logger:     p.Logger,
	}

	return wiring.RunnableGroup{
		Runnable: grpcServer,
	}, grpcServer
}

func (s *GRPCServer) Start() error {
	port := 50051 // get this from config
	s.logger.Bg().
		With(zap.Int("port", port)).
		Info("starting GRPC Server")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		s.logger.Bg().With(zap.Error(err)).Error("failed to listen")
		return err
	}

	if err := s.grpcServer.Serve(lis); err != nil {
		s.logger.Bg().With(zap.Error(err)).Error("failed to serve")
		return err
	}

	return nil
}

func (s *GRPCServer) Stop() error {
	s.grpcServer.GracefulStop()
	s.logger.Bg().Info("grpc server stopped")
	return nil
}
