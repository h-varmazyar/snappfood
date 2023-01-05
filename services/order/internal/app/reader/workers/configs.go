package workers

import "time"

type Configs struct {
	OrderWorkerInterval time.Duration
	OrderQueueKey       string
}
