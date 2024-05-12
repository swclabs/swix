package queue

import "swclabs/swipecore/internal/config"

var (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
	LowQueue      = "low"

	OrderQueue = "order"
	Purchase   = "purchase"
)

func init() {
	if config.StageStatus != "prod" {
		CriticalQueue = "critical_dev"
		DefaultQueue = "default_dev"
		LowQueue = "low_dev"
	}
}
