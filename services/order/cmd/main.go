package main

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/snappfood/pkg/serverext"
	"github.com/h-varmazyar/snappfood/services/order/configs"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/reader"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	logger := log.New()

	conf, err := loadConfigs()
	if err != nil {
		log.WithError(err).Panicf("failed to read configs")
		return
	}

	dbInstance, err := loadDB(ctx, conf.DB)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate databases")
	}

	if dbInstance == nil {
		logger.Panicf("nil db")
	}

	initializingApps(ctx, logger, dbInstance, conf)
}

func loadConfigs() (*Configs, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/app/configs")
	viper.AddConfigPath("/configs")
	if err := viper.ReadInConfig(); err != nil {
		localErr := viper.ReadConfig(bytes.NewBuffer(configs.DefaultConfig))
		if localErr != nil {
			return nil, localErr
		}
	}

	conf := new(Configs)
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func loadDB(ctx context.Context, configs *db.Configs) (*db.DB, error) {
	return db.NewDatabase(ctx, configs)
}

func initializingApps(ctx context.Context, logger *log.Logger, dbInstance *db.DB, configs *Configs) {
	var err error
	service := serverext.New(logger)
	var managerApp *manager.App
	managerApp, err = manager.NewApp(ctx, logger, dbInstance, configs.ManagerApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate reader app")
	}

	_, err = reader.NewApp(ctx, logger, dbInstance, configs.ReaderApp)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate reader app")
	}

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		managerApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Serve(configs.HttpPort, func(listener net.Listener) error {
		router := gin.Default()
		api := router.Group("/api")
		managerApp.Controller.RegisterRoutes(api)
		return http.Serve(listener, router)
	})

	service.Start(configs.ServiceName, configs.Version)
}
