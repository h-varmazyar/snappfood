package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/h-varmazyar/snappfood/pkg/errors"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type managerRedisRepository struct {
	redisDB *redis.Client
	logger  *log.Logger
}

func NewManagerRedisRepository(ctx context.Context, logger *log.Logger, db *redis.Client) (*managerRedisRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &managerRedisRepository{
		redisDB: db,
		logger:  logger,
	}, nil
}

func (db *managerRedisRepository) Create(ctx context.Context, order *entity.Order) error {
	err := db.redisDB.RPush(ctx, "order", order).Err()
	if err != nil {
		db.logger.WithError(err).Error("failed to push order into redis")
		return err
	}
	return nil
}
