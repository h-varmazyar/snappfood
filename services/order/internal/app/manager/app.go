package manager

import (
	"context"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/controller"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/repository"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/service"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/db"
	log "github.com/sirupsen/logrus"
)

type App struct {
	Service    *service.Service
	Controller *controller.Controller
}

func NewApp(ctx context.Context, logger *log.Logger, db *db.DB, configs *Configs) (*App, error) {
	repositoryInstance, err := repository.NewRepository(ctx, logger, configs.OrderQueue, db)
	if err != nil {
		return nil, err
	}

	managerService := service.NewService(ctx, logger, repositoryInstance)
	managerController := controller.NewController(logger, managerService)

	return &App{
		Service:    managerService,
		Controller: managerController,
	}, nil
}
