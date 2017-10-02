package cpu_usage

import (
	"github.com/gollector/gollector/logger"

	gm "github.com/gollector/gollector_metrics"
)

func GetMetric(params interface{}, log *logger.Logger) interface{} {
	results, err := gm.CPUUsage()

	if err != nil {
		log.Log("crit", err.Error())
		return nil
	}

	return results
}
