package service

import (
	"context"
	orderApi "github.com/h-varmazyar/snappfood/services/order/api/proto"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/repository"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	logger *log.Logger
	db     repository.ManagerRepository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, db repository.ManagerRepository) *Service {
	if GrpcService == nil {
		GrpcService = &Service{
			logger: logger,
			db:     db,
		}
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	orderApi.RegisterManagerServiceServer(server, s)
}

func (s *Service) CreateOrder(ctx context.Context, req *orderApi.ManagerCreateOrderReq) (*orderApi.Void, error) {
	order := &entity.Order{
		OrderID: req.OrderID,
		Price:   req.Price,
		Title:   req.Title,
	}

	err := s.db.Create(ctx, order)
	if err != nil {
		s.logger.WithError(err).Error("failed to create order")
		return nil, err
	}
	return new(orderApi.Void), nil
}
