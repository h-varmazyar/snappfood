package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/h-varmazyar/snappfood/pkg/errors"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type managerRedisRepository struct {
	redisDB *redis.Client
	logger  *log.Logger
	queue   string
}

func NewManagerRedisRepository(ctx context.Context, logger *log.Logger, queue string, db *redis.Client) (*managerRedisRepository, error) {
	if db == nil {
		return nil, errors.New(ctx, codes.Internal).AddDetailF("invalid db instance")
	}
	return &managerRedisRepository{
		redisDB: db,
		logger:  logger,
		queue:   queue,
	}, nil
}

func (db *managerRedisRepository) Create(ctx context.Context, order *entity.Order) error {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		db.logger.WithError(err).Error("cannot marshal order")
	}
	if err = db.redisDB.RPush(ctx, db.queue, string(orderBytes)).Err(); err != nil {
		db.logger.WithError(err).Error("failed to push order into redis")
		return err
	}
	return nil
}
