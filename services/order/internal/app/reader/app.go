package reader

import (
	"context"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/reader/repository"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/reader/workers"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	OrderWorker *workers.OrderWorker
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	orderWorker, err := workers.NewOrderWorker(ctx, logger, configs.WorkerConfigs, repositoryInstance, db.RedisDB)
	if err != nil {
		logger.WithError(err).Error("failed to initialize order worker")
		return nil, err
	}

	orderWorker.Start()

	return &App{
		OrderWorker: orderWorker,
	}, nil
}
