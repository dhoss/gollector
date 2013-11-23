package query

import (
	"fmt"
	"logger"
	"sync"
	"time"
	"types"
)

var rwmutex sync.RWMutex
var PluginResults types.PluginResultCollection

func GetResult(name string) types.PluginResult {
	return PluginResults[name]
}

func GetResults() types.PluginResultCollection {
	return PluginResults
}

func ResultPoller(config types.CirconusConfig, log *logger.Logger) {
	log.Log("info", "Starting Result Poller")
	interval_duration := time.Second * time.Duration(config.PollInterval)

	for {
		start := time.Now()
		AllResults(config, log)
		duration := time.Now().Sub(start)

		if duration < interval_duration {
			time.Sleep(interval_duration - duration)
		}
	}
}

func AllResults(config types.CirconusConfig, log *logger.Logger) {
	rwmutex.Lock()
	PluginResults = AllPlugins(config, log)
	rwmutex.Unlock()
}

func Plugin(name string, config types.CirconusConfig, log *logger.Logger) types.PluginResult {
	log.Log("debug", fmt.Sprintf("Plugin %s Requested", name))

	item, ok := config.Plugins[name]

	if ok {
		t, ok := types.Plugins[item.Type]

		if ok {
			log.Log("debug", fmt.Sprintf("Plugin %s exists, running", name))
			return t(item.Params, log)
		}
	}

	return nil
}

func AllPlugins(config types.CirconusConfig, log *logger.Logger) types.PluginResultCollection {
	retval := make(types.PluginResultCollection)

	log.Log("debug", "Querying All Plugins")

	for key, _ := range config.Plugins {
		retval[key] = Plugin(key, config, log)
	}

	log.Log("debug", "Done Querying All Plugins")

	return retval
}
