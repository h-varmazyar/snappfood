package workers

import "time"

type Configs struct {
	OrderWorkerInterval time.Duration `yaml:"order_worker_interval" mapstructure:"order_worker_interval"`
	OrderQueueKey       string        `yaml:"order_queue_key" mapstructure:"order_queue_key"`
}
