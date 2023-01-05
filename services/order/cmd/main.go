package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/snappfood/pkg/serverext"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	conf := loadConfigs()
	logger := log.New()
	dbInstance, err := loadDB(ctx, conf.DB)
	if err != nil {
		logger.WithError(err).Panicf("failed to initiate databases")
	}

	initializingApps(ctx, logger, dbInstance, conf)
}

func loadConfigs() *Configs {
	configs := new(Configs)
	confBytes, err := ioutil.ReadFile("../configs/local.yaml")
	if err != nil {
		log.WithError(err).Fatal("can not load yaml file")
	}
	if err = yaml.Unmarshal(confBytes, configs); err != nil {
		log.WithError(err).Fatal("can not unmarshal yaml file")
	}
	return configs
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

	service.Serve(configs.GRPCPort, func(lst net.Listener) error {
		server := grpc.NewServer()
		managerApp.Service.RegisterServer(server)
		return server.Serve(lst)
	})

	service.Serve(configs.HttpPort, func(listener net.Listener) error {
		router := gin.Default()
		managerApp.Controller.RegisterRoutes(router)
		return http.Serve(listener, router)
	})

	service.Start(configs.ServiceName, configs.Version)
}

//func registerHandlers(server *serverext.Server, port uint16, apps ) {
//	server.Serve(port, func(listener net.Listener) error {
//		router := gin.Default()
//
//		vouchers.NewHandler(configs.VouchersConfigs, logger).RegisterRoutes(router)
//
//		return http.Serve(listener, router)
//	})
//}
//
//func registerServices(server *serverext.Server, port netext.Port) {
//	server.Serve(port, func(listener net.Listener) error {
//		router := gin.Default()
//
//		vouchers.NewHandler(configs.VouchersConfigs, logger).RegisterRoutes(router)
//
//		return http.Serve(listener, router)
//	})
//}
