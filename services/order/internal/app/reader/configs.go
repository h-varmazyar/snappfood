package reader

import (
	"github.com/h-varmazyar/snappfood/services/order/internal/app/reader/workers"
)

type Configs struct {
	WorkerConfigs *workers.Configs `yaml:"worker_configs" mapstructure:"worker_configs"`
}
