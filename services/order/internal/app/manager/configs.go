package manager

import (
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/controller"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/service"
)

type Configs struct {
	ServiceConfigs    *service.Configs    `yaml:"service_configs"`
	ControllerConfigs *controller.Configs `yaml:"controller_configs"`
	//	BufferConfigs  *buffer.Configs  `yaml:"buffer_configs"`
	//	WorkerConfigs  *workers.Configs `yaml:"worker_configs"`
}
