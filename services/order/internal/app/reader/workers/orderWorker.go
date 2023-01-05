package workers

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/reader/repository"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type OrderWorker struct {
	db         repository.ReaderRepository
	configs    *Configs
	queue      *redis.Client
	ctx        context.Context
	ticker     *time.Ticker
	logger     *log.Logger
	cancelFunc context.CancelFunc
}

func NewOrderWorker(ctx context.Context, logger *log.Logger, configs *Configs, db repository.ReaderRepository, redisClient *redis.Client) (*OrderWorker, error) {
	workerCtx, cancelFunc := context.WithCancel(ctx)
	return &OrderWorker{
		db:         db,
		ctx:        workerCtx,
		configs:    configs,
		queue:      redisClient,
		logger:     logger,
		cancelFunc: cancelFunc,
	}, nil
}

func (w *OrderWorker) Start() {
	w.ticker = time.NewTicker(w.configs.OrderWorkerInterval)
	go w.do()
}

func (w *OrderWorker) do() {
	for {
		select {
		case <-w.ctx.Done():
			w.ticker.Stop()
			return
		case <-w.ticker.C:
			result, err := w.queue.BLPop(w.ctx, 0, w.configs.OrderQueueKey).Result()
			if err != nil {
				w.logger.WithError(err).Errorf("failed to pop fron redis")
				continue
			}
			w.handle(result)
		}
	}
}

func (w *OrderWorker) handle(data []string) {
	order := new(entity.Order)

	if err := json.NewDecoder(strings.NewReader(data[1])).Decode(&order); err != nil {
		w.logger.WithError(err).Error("failed to decode order")
		return
	}

	if err := w.db.Create(w.ctx, order); err != nil {
		w.logger.WithError(err).Errorf("failed to create order: %v", order)
	}
}
