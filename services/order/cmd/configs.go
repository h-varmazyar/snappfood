package main

import (
	"github.com/h-varmazyar/snappfood/pkg/netext"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/reader"
	"github.com/h-varmazyar/snappfood/services/order/internal/pkg/db"
)

type Configs struct {
	ServiceName string           `yaml:"service_name"`
	Version     string           `yaml:"version"`
	GRPCPort    netext.Port      `yaml:"grpc_port"`
	HttpPort    netext.Port      `yaml:"http_port"`
	DB          *db.Configs      `yaml:"db"`
	ManagerApp  *manager.Configs `yaml:"manager_app"`
	ReaderApp   *reader.Configs  `yaml:"reader_app"`
}
