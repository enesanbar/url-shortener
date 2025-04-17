package mappings

import (
	"context"

	"github.com/enesanbar/go-service/log"
	pb "github.com/enesanbar/proto-sdk-go/urlshortener/v1"
	"github.com/enesanbar/url-shortener/internal/usecase/mapping/create"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MappingsServerParams struct {
	fx.In

	Logger        log.Factory
	CreateUsecase create.Service `name:"producer"`
	DeleteUsecase create.Service `name:"producer"`
	UpdateUsecase create.Service `name:"producer"`
	GRPCServer    *GRPCServer
}

type MappingsServer struct {
	pb.UnimplementedUrlShortenerServiceServer

	Logger        log.Factory
	CreateUsecase create.Service `name:"producer"`
	DeleteUsecase create.Service `name:"producer"`
	UpdateUsecase create.Service `name:"producer"`
	GRPCServer    *GRPCServer
}

func NewMappingsServer(p MappingsServerParams) *MappingsServer {
	server := &MappingsServer{
		Logger:        p.Logger,
		CreateUsecase: p.CreateUsecase,
		DeleteUsecase: p.DeleteUsecase,
		UpdateUsecase: p.UpdateUsecase,
		GRPCServer:    p.GRPCServer,
	}
	pb.RegisterUrlShortenerServiceServer(p.GRPCServer.grpcServer, server)
	p.Logger.Bg().Info("mappings server registered in grpc server")
	return server
}

func (s *MappingsServer) CreateMapping(ctx context.Context, in *pb.CreateMappingRequest) (*pb.SingleMappingResponse, error) {
	s.Logger.For(ctx).Info("CreateMapping called")
	result, err := s.CreateUsecase.Execute(ctx, &create.Request{
		Code:      in.Code,
		URL:       in.Url,
		ExpiresAt: in.ExpiresAt.String(),
	})
	if err != nil {
		s.Logger.For(ctx).With(zap.Error(err)).Error("CreateMapping failed")
		return nil, err
	}

	s.Logger.For(ctx).Info("CreateMapping succeeded")
	return &pb.SingleMappingResponse{
		Response: &pb.ApiResponse{},
		Data: &pb.UrlMapping{
			Code:      result.Code,
			Url:       result.URL,
			ExpiresAt: timestamppb.New(*result.ExpiresAt),
			CreatedAt: timestamppb.New(result.CreatedAt),
			UpdatedAt: timestamppb.New(result.UpdatedAt),
		},
	}, nil
}

func (s *MappingsServer) GetMapping(ctx context.Context, in *pb.GetMappingRequest) (*pb.SingleMappingResponse, error) {
	return nil, nil
}

func (s *MappingsServer) GetMappings(ctx context.Context, in *pb.GetMappingsRequest) (*pb.MappingsListResponse, error) {
	return nil, nil
}

func (s *MappingsServer) UpdateMapping(ctx context.Context, in *pb.PatchMappingRequest) (*pb.SingleMappingResponse, error) {
	return nil, nil
}

func (s *MappingsServer) DeleteMapping(ctx context.Context, in *pb.DeleteMappingRequest) (*emptypb.Empty, error) {
	return nil, nil
}
