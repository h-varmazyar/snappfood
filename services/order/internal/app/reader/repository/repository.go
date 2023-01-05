package repository

import (
	"context"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/db"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
)

type ReaderRepository interface {
	Create(ctx context.Context, candle *entity.Order) error
}

func NewRepository(ctx context.Context, logger *log.Logger, db *db.DB) (ReaderRepository, error) {
	return NewReaderMySqlRepository(ctx, logger, db.MySqlDB)
}
