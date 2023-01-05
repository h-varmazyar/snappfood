package repository

import (
	"context"
	"github.com/h-varmazyar/snappfood/pkg/errors"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

type readerMySqlRepository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewReaderMySqlRepository(ctx context.Context, logger *log.Logger, db *gorm.DB) (*readerMySqlRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &readerMySqlRepository{
		db:     db,
		logger: logger,
	}, nil
}

func (r *readerMySqlRepository) Create(ctx context.Context, order *entity.Order) error {
	err := r.db.Model(new(entity.Order)).Create(order).Error
	if err != nil {
		r.logger.WithError(err).Error("failed to create new order")
		return err
	}
	return nil
}
