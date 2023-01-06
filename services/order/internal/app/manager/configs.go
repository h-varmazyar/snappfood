package manager

import (
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/controller"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/service"
)

type Configs struct {
	ServiceConfigs    *service.Configs    `yaml:"service_configs" mapstructure:"service_configs"`
	ControllerConfigs *controller.Configs `yaml:"controller_configs" mapstructure:"controller_configs"`
}
